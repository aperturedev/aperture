package httpecho

import (
	"github.com/aperturedev/aperture/internal/app"
	"github.com/aperturedev/aperture/internal/hyper"
	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
)

func RegisterDashboardServer(e *echo.Echo, ps *app.ProjectionService) {
	s := DashboardServer{}

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/projections")
	})

	e.GET("/projections", s.projections)
	e.GET("/projection/:id", s.projection)
}

type DashboardServer struct {
	svc *app.ProjectionService
	p   []app.Projection
}

func (s *DashboardServer) projections(c echo.Context) error {
	p, err := s.svc.AppProjections(c.Request().Context(), app.ProjectionsReq{
		AppID: "foo",
	})
	if err != nil {
		return err
	}

	s.p = p

	data := hyper.NewViewData().WithVal("projections", p)

	return c.Render(http.StatusOK, "projections", data)
}

func (s *DashboardServer) projection(c echo.Context) error {
	var pr app.Projection

	id := c.Param("id")

	for i, p := range s.p {
		if p.ID == id {
			r := uint64(rand.Intn(15))
			s.p[i].CurrentOffset += r

			if r%2 == 0 {
				s.p[i].State = app.ProjectionFailingState
			} else {
				s.p[i].State = app.ProjectionRunningState
			}

			pr = p
		}
	}

	data := hyper.NewViewData().With(pr)

	return c.Render(http.StatusOK, "projection", data)
}
