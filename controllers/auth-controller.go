package controllers

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/manjurulhoque/golang-job-portal/handlers"
	"github.com/manjurulhoque/golang-job-portal/models"
	"github.com/manjurulhoque/golang-job-portal/utils"
	"log"
	"net/http"
	"time"
)

type LoginResponse struct {
	token string
}

var (
	Secret     = "secret"
	ExpireTime = 3600
)

const (
	ErrorReason_ServerBusy = "server_busy"
	ErrorReason_ReLogin    = "Relogin"
)

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid_data", "message": err.Error()})
		c.Abort()
		return
	}

	errs := utils.TranslateError(user)
	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs})
		return
	}

	err := handlers.Register(&user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, user.UserOutput())
	}
}

func Login(c *gin.Context) {

	var user models.LoginInput
	if c.BindJSON(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid_data", "message": "Invalid data"})
		c.Abort()
		return
	}
	errs := utils.TranslateError(user)
	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errs})
		return
	}

	err := handlers.Login(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	claims := &handlers.JWTClaims{
		Email:  user.Email,
		UserId: user.ID,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	signedToken, err := getToken(claims)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"access": signedToken})
}

//http://localhost:9090/verify/eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjA1MTIyMTAsImlhdCI6MTU2MDUwODYxMCwidXNlcl9pZCI6MSwicGFzc3dvcmQiOiIxMjM0NTYiLCJ1c2VybmFtZSI6ImRvbmciLCJmdWxsX25hbWUiOiJkb25nIiwicGVybWlzc2lvbnMiOltdfQ.Esh1Zge0vO1BAW1GeR5wurWP3H1jUIaMf3tcSaUwkzA
func Verify(c *gin.Context) {
	strToken := c.Param("token")
	matched, err := utils.RegexpToken(strToken)
	if err != nil || !matched {
		c.String(http.StatusNotFound, "token invalid")
		return
	}
	claim, err := VerifyAction(strToken)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("%s,%s", "verify success", claim.Email))
}

//http://localhost:9090/refresh/eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjA1MTIyNDMsImlhdCI6MTU2MDUwODYxMCwidXNlcl9pZCI6MSwicGFzc3dvcmQiOiIxMjM0NTYiLCJ1c2VybmFtZSI6ImRvbmciLCJmdWxsX25hbWUiOiJkb25nIiwicGVybWlzc2lvbnMiOltdfQ.Xkb_J8MWXkwGUcBF9bpp2Ccxp8nFPtRzFzOBeboHmg0
func Refresh(c *gin.Context) {
	strToken := c.Param("token")
	claims, err := VerifyAction(strToken)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	claims.ExpiresAt = time.Now().Unix() + (claims.ExpiresAt - claims.IssuedAt)
	signedToken, err := getToken(claims)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, signedToken)
}

func VerifyAction(strToken string) (*handlers.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(strToken, &handlers.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})

	if err != nil {
		log.Print(err.Error())
		return nil, errors.New("unauthorized")
	}
	claims, ok := token.Claims.(*handlers.JWTClaims)
	if !ok {
		return nil, errors.New("unauthorized")
	}
	if err := token.Claims.Valid(); err != nil {
		log.Print(err.Error())
		return nil, errors.New("unauthorized")
	}
	return claims, nil
}

func getToken(claims *handlers.JWTClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", errors.New(ErrorReason_ServerBusy)
	}
	return signedToken, nil
}

func CurrentUserTodos(c *gin.Context) {
	claims, _ := c.Get("Claims")

	claims2 := claims.(*handlers.JWTClaims)

	user := handlers.FindUserById(claims2.UserId)

	c.JSON(http.StatusOK, gin.H{"user": user})
}
