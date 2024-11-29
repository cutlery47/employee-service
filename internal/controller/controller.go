package controller

import repo "github.com/cutlery47/auth-service/internal/repository"

type Controller struct {
	repo *repo.Repository
}

func NewController(repo *repo.Repository) *Controller {
	return &Controller{
		repo: repo,
	}
}

func (c *Controller) handleGet() {
	return
}
