package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/manjurulhoque/golang-job-portal/utils"

	//_ "github.com/go-ozzo/ozzo-validation/v4"
	//_ "github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/manjurulhoque/golang-job-portal/models"
	"net/http"
)

var validate *validator.Validate = validator.New()

func CreateJob(c *gin.Context) {
	var job models.Job

	if err := c.BindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid_data", "message": err.Error()})
		return
	}

	fmt.Println(job)

	err := validate.Struct(job)

	errs := utils.TranslateError(err, validate)

	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid_data", "errors": errs})
		return
	}
}
