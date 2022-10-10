package models

import (
	"golang.org/x/crypto/bcrypt"
	"html"
	"strings"
	"time"
)

type User struct {
	BaseModel
	Email    string `gorm:"unique" json:"email" validate:"required,email,emailExists"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Jobs     []Job  `json:"jobs"`
	Role     string `gorm:"not null,type:enum('employee', 'employer');not null" json:"role" validate:"required,validRole"`
}

type UserJob struct {
	BaseModel
	Email string `gorm:"unique" json:"email" validate:"required,email,emailExists"`
	Name  string `json:"name" validate:"required"`
}

type UpdateUserProfile struct {
	Name string `json:"name" validate:"required"`
}

func (UserJob) TableName() string {
	return "users"
}

type RetrieveUser struct {
	BaseModel
	Email string `json:"email"`
	Name  string `json:"name"`
	Jobs  []Job  `json:"jobs"`
	Role  string `json:"role"`
}

type RegisterInput struct {
	Email    string `json:"email" validate:"required,email,emailExists"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `gorm:"not null,type:enum('employee', 'employer');not null" json:"role" validate:"required,validRole"`
}

type RegisterData struct {
	BaseModel
	Email    string `json:"email" validate:"required,email,emailExists"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `gorm:"not null,type:enum('employee', 'employer');not null" json:"role" validate:"required,validRole"`
}

type LoginData struct {
	BaseModel
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (RetrieveUser) TableName() string {
	return "users"
}

func (LoginData) TableName() string {
	return "users"
}

func (LoginInput) TableName() string {
	return "users"
}

func (u *RegisterData) TableName() string {
	return "users"
}

func (RegisterInput) TableName() string {
	return "users"
}

func (u *User) TableName() string {
	return "users"
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *RegisterData) UserOutput() User {
	return User{
		BaseModel: BaseModel{u.ID, u.CreatedAt, u.UpdatedAt, u.DeletedAt},
		Email:     u.Email,
		Name:      u.Name,
		Role:      u.Role,
	}
}

func (u *User) UserOutput() User {
	return User{
		BaseModel: BaseModel{u.ID, u.CreatedAt, u.UpdatedAt, u.DeletedAt},
		Email:     u.Email,
		Name:      u.Name,
	}
}

func (u *User) Prepare() {
	u.ID = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}
