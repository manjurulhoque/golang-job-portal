package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/manjurulhoque/golang-job-portal/config"
	"github.com/manjurulhoque/golang-job-portal/models"
	"github.com/manjurulhoque/golang-job-portal/utils"
	"net/http"
)

func AllTags(c *gin.Context) {
	var tags []models.Tag

	config.DB.Find(&tags)

	c.JSON(http.StatusOK, tags)
}

func CreateTag(c *gin.Context) {
	var tagInput models.Tag

	if err := c.ShouldBindJSON(&tagInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	errs := utils.TranslateError(tagInput)

	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errs})
		return
	}

	if err := config.DB.Create(&tagInput).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, tagInput)
	}
}
