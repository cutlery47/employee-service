package controller

import (
	repo "github.com/cutlery47/employee-service/internal/repository"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Controller struct {
	repo *repo.Repository
}

func NewController(repo *repo.Repository, e *echo.Echo) *Controller {

	e.GET("/ping", func(c echo.Context) error { return c.NoContent(200) })
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := e.Group("/api/v1")

	ctl := &Controller{
		repo: repo,
	}

	v1.GET("/employees", ctl.handleGet)
	v1.GET("/employee", ctl.handleGetMeta)

	return ctl
}

func (ctl *Controller) handleGet(c echo.Context) error {
	return nil
}

func (ctl *Controller) handleGetMeta(c echo.Context) error {
	return nil
}
