package model

import (
	"time"
)

type UserMeta struct {
	UserGeneral
	DateOfBirth time.Time
	Cellphone   string
	Email       string
	Address     string
	City        string
	Team        string
	Position    string
}

type UserGeneral struct {
	Id         int
	Name       string
	Surname    string
	Department string
	Role       string
}
