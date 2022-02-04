package handlers

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"github.com/manjurulhoque/golang-job-portal/config"
	"github.com/manjurulhoque/golang-job-portal/models"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type JWTClaims struct {
	jwt.StandardClaims
	Email  string `json:"email"`
	UserId uint   `json:"user_id"`
}

func Login(user *models.LoginInput) (err error) {

	//user := models.User{}
	previousPassword := user.Password

	if err := config.DB.Table("users").Where("email = ?", user.Email).First(&user).Error; err != nil {
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

func CheckUserExists(email string) (exists bool, err error) {
	r := config.DB.Table("users").Where("email = ?", email).Limit(1)

	if r.Error != nil {
		err = r.Error
		return false, err
	}
	userExists := r.RowsAffected > 0

	err = errors.New("a user is already exists with this email")

	return userExists, err
}

func FindUserByEmail(email string) (user models.RetrieveUser, exists bool) {
	result := config.DB.Table("users").Where("email = ?", email).Take(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		logrus.Info(result.Error.Error())
	}

	return user, result.RowsAffected > 0
}

func FindUserById(userId uint) (user models.RetrieveUser) {
	var userModel models.User
	err := config.DB.Table("users").Preload("Jobs").Find(&userModel, userId).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println(err.Error())
	}

	if err := copier.Copy(&user, &userModel); err != nil {
		logrus.Error(err)
	}

	return user
}
