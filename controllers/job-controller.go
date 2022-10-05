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

// AllJobs All unfilled jobs
// @Summary All unfilled jobs
// @Description All unfilled jobs
// @Tags jobs
// @Accept application/json
// @Produce json
// @Success 200
// @Router /jobs/ [get]
func AllJobs(c *gin.Context) {
	var jobs []models.Job

	config.DB.Preload("Tags").Preload("User").Where(models.Job{Filled: false}).Find(&jobs)
	//config.DB.Preload("User").Where(models.Job{Filled: false}).Find(&jobs)

	c.JSON(http.StatusOK, jobs)
}

// JobDetails Job details
// @Summary Job details
// @Description Job details
// @Tags jobs
// @Accept application/json
// @Produce json
// @Success 200
// @Router /jobs/ [get]
func JobDetails(c *gin.Context) {
	var job models.Job
	var jobId = c.Param("job_id")

	if err := config.DB.Where("id = ?", jobId).First(&job).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Record not found!"})
		return
	}
	if job.Tags == nil {
		job.Tags = []models.TagJob{}
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(job))
}

// CreateJob Create new job
// @Summary Create new job
// @Description Create new job as employer
// @Tags jobs
// @Accept application/json
// @Produce json
// @Success 200
// @Router /jobs/create [post]
func CreateJob(c *gin.Context) {
	var jobInput models.JobInput
	var newJob models.Job
	var tags []models.Tag

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
	for _, tagId := range jobInput.Tags {
		var tag models.Tag
		config.DB.Where("id = ?", tagId).First(&tag)
		if tag.ID != 0 {
			tags = append(tags, tag)
		}
	}

	if err := config.DB.Create(&newJob).Association("Tags").Append(tags).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusCreated, utils.SuccessResponse(newJob))
	}
}

// UpdateJob Update job
// @Summary Update job
// @Description Update job as employer
// @Tags jobs
// @Accept application/json
// @Produce json
// @Success 200
// @Router /jobs/ [post]
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

	c.JSON(http.StatusOK, utils.SuccessResponse(existingJob))
}

// ApplyToTheJob Apply for job
// @Summary Apply to the job
// @Description Apply to the job
// @Tags jobs, employee
// @Accept application/json
// @Produce json
// @Success 200
// @Router /jobs/:job_id/apply-job [post]
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
		c.JSON(http.StatusOK, utils.SuccessResponse(newApplicant))
	}
}

// AppliedJobs
// @Summary Get applied jobs
// @Description Get all applied jobs for current logged in employee
// @Tags jobs, employee
// @Accept application/json
// @Produce json
// @Success 200
// @Router /jobs/applied-jobs [get]
func AppliedJobs(c *gin.Context) {
	user, _ := utils.AuthorizedUser(c)

	var applicants []models.Applicant

	//config.DB.Raw("select id, status, comment, user_id, job_id from applicants where user_id = ?", user.ID).Scan(&applicants)

	config.DB.Preload("Job").Where(models.Applicant{UserId: user.ID}).Find(&applicants)
	c.JSON(http.StatusOK, utils.SuccessResponse(applicants))
}
