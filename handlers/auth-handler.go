package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/manjurulhoque/golang-job-portal/config"
	"github.com/manjurulhoque/golang-job-portal/models"
	"golang.org/x/crypto/bcrypt"
)

type JWTClaims struct {
	jwt.StandardClaims
	Email  string `json:"email"`
	UserId uint   `json:"user_id"`
}

func Login(user *models.User) (err error) {

	//user := models.User{}
	previousPassword := user.Password

	if err := config.DB.Where("email = ?", user.Email).First(&user).Error; err != nil {
		return err
	}

	err = models.VerifyPassword(user.Password, previousPassword)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return err
	}

	return nil
}

func Register(user *models.User) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	if err != nil {
		return err
	}

	if err := config.DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func FindUserById(userId uint) (user models.RetrieveUser) {
	config.DB.Table("users").Find(&user, userId)

	return user
}
