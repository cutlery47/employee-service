package controller

import (
	repo "github.com/cutlery47/employee-service/internal/repository"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Controller struct {
	repo *repo.Repository
}

// добавить контроллер с хинтами

func NewController(repo *repo.Repository, e *echo.Echo) *Controller {

	e.GET("/ping", func(c echo.Context) error { return c.NoContent(200) })
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := e.Group("/api/v1")

	ctl := &Controller{
		repo: repo,
	}

	v1.POST("/employees", ctl.handleGet)
	v1.POST("/employee", ctl.handleGetMeta)

	return ctl
}

// city string (optional)
// position string (opitonal)
// part string (optional)
// project (optional)
// unit (optional)
// fullname (optional)
// no args - error

func (ctl *Controller) handleGet(c echo.Context) error {
	// ctx := c.Request().Context()

	// params := c.QueryParams()

	// filters, err := model.FromURL(params)
	// if err != nil {
	// 	return handleError(err)
	// }

	// if !params.Has("limit") {
	// 	return echo.NewHTTPError(400, "pagination limit was not provided")
	// }

	// if !params.Has("offset") {
	// 	return echo.NewHTTPError(400, "pagination offset was not provided")
	// }

	// limit, err := strconv.Atoi(params.Get("limit"))
	// if err != nil {
	// 	return echo.NewHTTPError(400, "couldn't parse pagination limit")
	// }

	// offset, err := strconv.Atoi(params.Get("offset"))
	// if err != nil {
	// 	return echo.NewHTTPError(400, "couldn't parse pagination offset")
	// }

	// res, err := ctl.repo.Get(ctx, filters, limit, offset)
	// if err != nil {
	// 	return handleError(err)
	// }

	// return c.JSON(200, res)
	return nil
}

func (ctl *Controller) handleGetMeta(c echo.Context) error {
	// ctx := c.Request().Context()

	// id := c.QueryParam("id")
	// if id == "" {
	// 	return echo.NewHTTPError(400, "id was not provided")
	// }

	// intId, err := strconv.Atoi(id)
	// if err != nil {
	// 	return echo.NewHTTPError(400, "id was not resolved")
	// }

	// meta, err := ctl.repo.GetMeta(ctx, intId)
	// if err != nil {
	// 	return handleError(err)
	// }

	// return c.JSON(200, meta)
	return nil
}
