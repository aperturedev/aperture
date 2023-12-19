package httpecho

import (
	"github.com/aperturedev/aperture/internal/app"
	"github.com/aperturedev/aperture/internal/hyper"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterDashboardServer(e *echo.Echo) {
	s := DashboardServer{}

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/projections")
	})

	e.GET("/projections", s.projections)
}

type DashboardServer struct {
}

func (s *DashboardServer) projections(c echo.Context) error {
	data := hyper.NewViewData().WithVal("projections", []app.Projection{
		{
			Name:        "New Employees",
			Description: "Creates a read model of new employees in order to be used by internal hr people",
			State:       app.ProjectionRunningState,
		},
		{
			Name:        "Time Sheets",
			Description: "Creates a read model of new employees in order to be used by internal hr people",
			State:       app.ProjectionFailingState,
		},
		{
			Name:        "Absences and Sick Leaves",
			Description: "Creates a read model of new employees in order to be used by internal hr people",
			State:       app.ProjectionRunningState,
		},
	})

	return c.Render(http.StatusOK, "projections", data)
}
