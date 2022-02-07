package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/manjurulhoque/golang-job-portal/config"
	"github.com/manjurulhoque/golang-job-portal/handlers"
	"github.com/manjurulhoque/golang-job-portal/utils"
	"github.com/sirupsen/logrus"
	"strconv"

	//_ "github.com/go-ozzo/ozzo-validation/v4"
	//_ "github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/manjurulhoque/golang-job-portal/models"
	"net/http"
)

//var validate = validator.New()

func AllJobs(c *gin.Context) {
	var jobs []models.Job

	config.DB.Preload("Tags").Find(&jobs)

	c.JSON(http.StatusOK, jobs)
}

func CreateJob(c *gin.Context) {
	var jobInput models.JobInput
	var newJob models.Job

	if err := c.ShouldBindJSON(&jobInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, _ := utils.AuthorizedUser(c)

	errs := utils.TranslateError(jobInput)

	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errs})
		return
	}

	if err := copier.Copy(&newJob, &jobInput); err != nil {
		logrus.Error(err)
	}

	newJob.UserId = user.ID
	//newJob.Tags = jobInput.Tags

	if err := config.DB.Create(&newJob).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusCreated, newJob)
	}
}

func UpdateJob(c *gin.Context) {
	var jobInput models.JobInput
	var existingJob models.Job
	var jobId = c.Param("job_id")

	if err := c.ShouldBindJSON(&jobInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	errs := utils.TranslateError(jobInput)
	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errs})
		return
	}
	if err := config.DB.Where("id = ?", jobId).First(&existingJob).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Record not found!"})
		return
	}
	if !utils.RequesterIsJobOwner(c, &existingJob) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to access this resource"})
		return
	}
	logrus.Info(existingJob)

	if err := config.DB.Model(&existingJob).Updates(jobInput).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, existingJob)
}

// ApplyToTheJob Apply for job
func ApplyToTheJob(c *gin.Context) {
	var job models.Job
	jobId, _ := strconv.Atoi(c.Param("job_id"))

	if err := config.DB.Where("id = ?", jobId).First(&job).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Job not found!"})
		return
	}
	user, _ := utils.AuthorizedUser(c)
	alreadyApplied := handlers.AlreadyAppliedForTheJob(c, &job, &user)

	if alreadyApplied {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "You already applied for the job"})
		return
	}

	newApplicant := models.Applicant{
		JobId:  uint(jobId),
		UserId: user.ID,
	}

	if err := config.DB.Create(&newApplicant).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, newApplicant)
	}
}
