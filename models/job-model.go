package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
)

type Job struct {
	gorm.Model
	Title       string      `json:"title" validate:"required"`
	Description string      `gorm:"type:text" json:"description" validate:"required"`
	UserId      uint        `json:"user_id"`
	Salary      json.Number `json:"salary" validate:"required,integer"`
	Location    string      `json:"location" validate:"required"`

	User User `gorm:"foreignkey:UserId" json:"user"`
}

type JobInput struct {
	gorm.Model
	Title       string      `json:"title" validate:"required"`
	Description string      `json:"description" validate:"required"`
	Salary      json.Number `json:"salary" validate:"required,integer"`
	Location    string      `json:"location" validate:"required"`
}
