package hyper

import (
	"reflect"

	"github.com/labstack/echo/v4"
)

func NewViewData() ViewData {
	return ViewData{
		Data:   make(map[string]any),
		Errors: make(map[string]any),
	}
}

type ViewData struct {
	echo.Context
	Data   map[string]any
	Errors map[string]any
}

func (vd ViewData) With(val any) ViewData {
	return vd.WithVal(reflect.TypeOf(val).Name(), val)
}

func (vd ViewData) WithVal(key string, val any) ViewData {
	vd.Data[key] = val

	return vd
}

func (vd ViewData) WithError(key string, val any) ViewData {
	vd.Errors[key] = val

	return vd
}
