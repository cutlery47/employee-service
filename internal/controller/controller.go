package controller

import (
	"strconv"

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
	ctx := c.Request().Context()

	id := c.QueryParam("id")
	if id == "" {
		return echo.NewHTTPError(400, "id was not provided")
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		return echo.ErrInternalServerError
	}

	meta, err := ctl.repo.GetMeta(ctx, intId)
	if err != nil {
		return ctl.handleError(err)
	}

	return c.JSON(200, meta)
}

func (ctl *Controller) handleError(err error) *echo.HTTPError {
	return echo.ErrInternalServerError
}
