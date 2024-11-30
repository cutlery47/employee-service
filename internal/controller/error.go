package controller

import (
	"github.com/cutlery47/employee-service/internal/repository"
	"github.com/labstack/echo/v4"
)

var errMap = map[error]*echo.HTTPError{
	repository.ErrUserNotFound:    echo.ErrNotFound,
	repository.ErrWrongDateFormat: echo.ErrBadRequest,
}

func handleError(err error) *echo.HTTPError {
	if httpErr, ok := errMap[err]; ok {
		httpErr.Message = err.Error()
		return httpErr
	}

	return echo.ErrInternalServerError
}
