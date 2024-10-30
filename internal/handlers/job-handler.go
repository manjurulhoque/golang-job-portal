package handlers

import (
	"context"
	"github.com/manjurulhoque/golang-job-portal/internal/config"
	"github.com/manjurulhoque/golang-job-portal/internal/models"
	"log/slog"
)

// AlreadyAppliedForTheJob Check if user already applied for the job
func AlreadyAppliedForTheJob(ctx context.Context, job *models.Job, user *models.RetrieveUser) bool {

	var alreadyAppliedApplicant models.Applicant
	result := config.DB.Where("user_id = ? AND job_id = ?", user.ID, job.ID).Find(&alreadyAppliedApplicant)
	if result.Error != nil {
		slog.Error("Error while checking if user already applied for the job", "error", result.Error.Error())
		return false
	}

	return result.RowsAffected > 0
}
