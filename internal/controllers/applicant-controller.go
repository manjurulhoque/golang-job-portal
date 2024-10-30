package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/manjurulhoque/golang-job-portal/internal/config"
	"github.com/manjurulhoque/golang-job-portal/internal/models"
	"github.com/manjurulhoque/golang-job-portal/pkg/utils"
	"net/http"
)

// ApplicantsForEmployer godoc
// @Summary Applicants for employer.
// @Description Applicants for employer.
// @Tags applicants
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /jobs/applicants [get]
func ApplicantsForEmployer(c *gin.Context) {
	//user, _ := utils.AuthorizedUser(c)

	type Result struct {
		models.BaseModel
		Comment string `json:"comment"`
		Status  int    `json:"status"`
		UserId  uint   `json:"user_id"`
		JobId   uint   `json:"job_id"`
	}

	var applicants []Result

	config.DB.Raw("select * from applicants inner join jobs on applicants.job_id=jobs.id").Scan(&applicants)
	c.JSON(http.StatusOK, utils.SuccessResponse(applicants))
}
