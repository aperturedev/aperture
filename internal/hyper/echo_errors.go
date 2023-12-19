package hyper

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// NewErrorHandler creates custom http error handler
func NewErrorHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		var appErr AppError

		if errors.As(err, &appErr) {
			_ = c.Render(http.StatusOK, appErr.tpl, appErr.vd)

			return
		}

		_ = c.Render(http.StatusOK, "server_error", NewViewData().WithError("ServerError", err.Error()))
	}
}

func NewAppError() AppError {
	return AppError{
		tpl: "server_error",
	}
}

type AppError struct {
	err error
	tpl string
	vd  ViewData
}

func (ve AppError) Error() string {
	return "some error"
}

func (ve AppError) WithTpl(name string) AppError {
	ve.tpl = name

	return ve
}

func (ve AppError) WithData(data ViewData) AppError {
	ve.vd = data

	return ve
}

func (ve AppError) WithErr(err error) AppError {
	ve.err = err

	return ve
}
