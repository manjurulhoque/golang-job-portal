package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/manjurulhoque/golang-job-portal/config"
	"github.com/manjurulhoque/golang-job-portal/models"
	"github.com/sirupsen/logrus"
)

// AlreadyAppliedForTheJob Check if user already applied for the job
func AlreadyAppliedForTheJob(c *gin.Context, job *models.Job, user *models.RetrieveUser) bool {

	var alreadyAppliedApplicant models.Applicant
	result := config.DB.Where("user_id = ? AND job_id = ?", user.ID, job.ID).Find(&alreadyAppliedApplicant)
	if result.Error != nil {
		logrus.Error(result.Error.Error())
		return false
	}

	return result.RowsAffected > 0
}
