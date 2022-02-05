package models

import (
	"encoding/json"
)

type Job struct {
	BaseModel
	Title       string `json:"title" validate:"required"`
	Description string `gorm:"type:text" json:"description" validate:"required"`
	UserId      uint   `json:"user_id"`
	Salary      int    `json:"salary" validate:"required,integer"`
	Location    string `json:"location" validate:"required"`
	Filled      bool   `gorm:"default:false" json:"filled"`

	User User `gorm:"foreignkey:UserId" json:"user"`
}

type JobInput struct {
	BaseModel
	Title       string      `json:"title" validate:"required"`
	Description string      `json:"description" validate:"required"`
	Salary      json.Number `json:"salary" validate:"required,integer"`
	Location    string      `json:"location" validate:"required"`
}
