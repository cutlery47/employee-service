package controller

import (
	"encoding/json"

	"github.com/cutlery47/employee-service/internal/model"
	repo "github.com/cutlery47/employee-service/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Controller struct {
	repo *repo.Repository
	h    *errHandler
}

func NewController(repo *repo.Repository, e *echo.Echo, errLog, infoLog *logrus.Logger) {
	ctl := &Controller{
		repo: repo,
		h: &errHandler{
			errLog: errLog,
		},
	}

	e.GET("/ping", func(c echo.Context) error { return c.NoContent(200) })
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := e.Group("/api/v1", requestLoggerMiddleware(infoLog))
	{
		v1.POST("/employees", ctl.GetEmployee)
		v1.POST("/employee", ctl.GetBaseEmpoyees)
	}
}

// /api/v1/employee (POST)
func (ctl *Controller) GetEmployee(c echo.Context) error {
	ctx := c.Request().Context()

	body := c.Request().Body

	request := struct {
		Id int `json:"id,omitempty"`
	}{}

	decoder := json.NewDecoder(body)
	err := decoder.Decode(&request)
	if err != nil {
		ctl.h.handleError(err)
	}

	res, err := ctl.repo.GetEmployee(ctx, request.Id)
	if err != nil {
		return ctl.h.handleError(err)
	}

	return c.JSON(200, res)
}

// api/v1/employees
func (ctl *Controller) GetBaseEmpoyees(c echo.Context) error {
	ctx := c.Request().Context()

	body := c.Request().Body

	request := model.GetBaseEmployeesRequest{}

	decoder := json.NewDecoder(body)
	err := decoder.Decode(&request)
	if err != nil {
		ctl.h.handleError(err)
	}

	employees, err := ctl.repo.GetBaseEmployees(ctx, request)
	if err != nil {
		return ctl.h.handleError(err)
	}

	return c.JSON(200, employees)
}
