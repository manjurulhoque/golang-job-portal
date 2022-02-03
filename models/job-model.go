package models

import "github.com/jinzhu/gorm"

type Job struct {
	gorm.Model
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

//UserId      int
//User        User