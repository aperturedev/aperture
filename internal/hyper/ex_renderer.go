package hyper

import (
	"fmt"
	"io"

	"github.com/Masterminds/sprig/v3"
	"github.com/dannyvankooten/extemplate"
	"github.com/labstack/echo/v4"
)

func NewRenderer(viewPath string) *Renderer {
	xt := extemplate.New().
		Funcs(sprig.FuncMap())

	err := xt.ParseDir(fmt.Sprintf("%s/", viewPath), []string{".tpl"})
	if err != nil {
		panic(err)
	}

	return &Renderer{
		tpl: xt,
	}
}

type Renderer struct {
	tpl *extemplate.Extemplate
}

func (r *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	viewData, ok := data.(ViewData)
	if ok {
		viewData.Context = c
	}

	return r.tpl.ExecuteTemplate(w, fmt.Sprintf("%s.go.tpl", name), viewData)
}
