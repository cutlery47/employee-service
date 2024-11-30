package controller

import (
	"encoding/json"

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

	v1.POST("/employees", ctl.Get)
	v1.POST("/employee", ctl.GetMeta)

	return ctl
}

// city string (optional)
// position string (opitonal)
// part string (optional)
// project (optional)
// unit (optional)
// fullname (optional)
// no args - error

// /api/v1/employee (POST)
func (ctl *Controller) Get(c echo.Context) error {
	ctx := c.Request().Context()

	body := c.Request().Body

	request := struct {
		id int
	}{}

	decoder := json.NewDecoder(body)
	err := decoder.Decode(&request)

	if err != nil {
		panic(err)
	}

	res, err := ctl.repo.GetEmployee(ctx, request.id)
	if err != nil {
		return handleError(err)
	}

	return c.JSON(200, res)
}

// api/v1/employees
func (ctl *Controller) GetMeta(c echo.Context) error {
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
