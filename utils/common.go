package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/manjurulhoque/golang-job-portal/models"
	"net/http"
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

func RequesterIsJobOwner(c *gin.Context, job *models.Job) bool {

	user, err := AuthorizedUser(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return false
	}

	return job.UserId == user.ID
}
