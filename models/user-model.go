package models

import (
	"errors"
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
	Role     string `gorm:"type:enum('employee', 'employer');not null" json:"role" validate:"required,validRole"`
}

type RetrieveUser struct {
	BaseModel
	Email string `json:"email"`
	Name  string `json:"name"`
	Jobs  []Job  `json:"jobs"`
	Role  string `json:"role"`
}

type LoginInput struct {
	BaseModel
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (User) TableName() string {
	return "users"
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
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

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Password == "" {
			return errors.New("password is required")
		}
		return nil
	case "login":
		if u.Password == "" {
			return errors.New("password is required")
		}
		if u.Email == "" {
			return errors.New("email is required")
		}
		return nil

	case "register":
		if u.Email == "" {
			return errors.New("email is required")
		}
		if u.Name == "" {
			return errors.New("name is required")
		}
		if u.Password == "" {
			return errors.New("password is required")
		}
		return nil

	default:
		return nil
	}
}
