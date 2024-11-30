package controller

import (
	"encoding/json"
	"fmt"

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

func (ctl *Controller) GetHint(c echo.Context) error {
	ctx := c.Request().Context()
	body := c.Request().Body

	request := struct {
		City     string `json:"city_search_term,omitempty"`
		Position string `json:"position_search_term,omitempty"`
		Project  string `json:"project_search_term,omitempty"`
		Role     string `json:"role_search_term,omitempty"`
		Unit     string `json:"unit_search_term,omitempty"`
	}{}

	decoder := json.NewDecoder(body)
	err := decoder.Decode(&request)
	if err != nil {
		ctl.h.handleError(err)
	}

	var res interface{}
	if request.City != "" {
		res, err = ctl.repo.GetHints(ctx, "city", request.City)
	} else if request.Position != "" {
		res, err = ctl.repo.GetHints(ctx, "position", request.Position)
	} else if request.Project != "" {
		res, err = ctl.repo.GetHints(ctx, "project", request.Project)
	} else if request.Role != "" {
		res, err = ctl.repo.GetHints(ctx, "role", request.Role)
	} else if request.Unit != "" {
		res, err = ctl.repo.GetHints(ctx, "unit", request.Unit)
	} else {
		return fmt.Errorf("error: no args")
	}
	if err != nil {
		return ctl.h.handleError(err)
	}

	return c.JSON(200, res)
}
