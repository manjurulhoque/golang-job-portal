package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/manjurulhoque/golang-job-portal/internal/config"
	"github.com/manjurulhoque/golang-job-portal/internal/models"
	"net/http"
)

func UpdateUserProfile(c *gin.Context, newUserData *models.UpdateUserProfile, user *models.RetrieveUser) (err error) {

	if err := config.DB.Model(&user).Updates(newUserData).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return nil
	}

	return nil
}
