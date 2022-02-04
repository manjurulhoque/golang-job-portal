package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/manjurulhoque/golang-job-portal/config"
	"github.com/manjurulhoque/golang-job-portal/utils"

	//_ "github.com/go-ozzo/ozzo-validation/v4"
	//_ "github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/manjurulhoque/golang-job-portal/models"
	"net/http"
)

//var validate = validator.New()

func CreateJob(c *gin.Context) {
	var job models.Job

	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	errs := utils.TranslateError(job)

	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errs})
		return
	}
	if err := config.DB.Create(&job).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, job)
	}
}
