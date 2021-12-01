package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"html"
	"strings"
	"time"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique" json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RetrieveUser struct {
	gorm.Model
	Email string `json:"email"`
	Name  string `json:"name"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "users"
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
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
