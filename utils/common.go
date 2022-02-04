package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/manjurulhoque/golang-job-portal/models"
)

// AuthorizedUser .
func AuthorizedUser(c *gin.Context) (user models.RetrieveUser, err error) {
	authorizedUser, ok := c.Get("AuthorizedUser")
	if !ok {
		err = errors.New("no 'AuthorizedUser'")
		return
	}
	user, ok = authorizedUser.(models.RetrieveUser)
	if !ok {
		err = errors.New("'AuthorizedUser' not 'models.User' type")
		return
	}
	return user, nil
}
