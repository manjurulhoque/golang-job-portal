package models

import "github.com/jinzhu/gorm"

type Job struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      int
	User        User
}
