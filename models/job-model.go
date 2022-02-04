package models

import "github.com/jinzhu/gorm"

type Job struct {
	gorm.Model
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	UserId      uint   `json:"user_id"`
	User        User   `gorm:"foreignkey:UserId" json:"user"`
}

type JobInput struct {
	gorm.Model
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}
