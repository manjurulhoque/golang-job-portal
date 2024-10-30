package routes

import (
	"github.com/gin-gonic/gin"
	controllers2 "github.com/manjurulhoque/golang-job-portal/internal/controllers"
	"github.com/manjurulhoque/golang-job-portal/internal/middlewares"
)

func (r routes) addJobRoutes(rg *gin.RouterGroup) {
	v1 := rg.Group("/jobs")

	v1.GET("/", controllers2.AllJobs)
	v1.GET("/:job_id", controllers2.JobDetails)

	v1.Use(middlewares.AuthMiddleware())
	{
		v1.GET("/user", controllers2.CurrentUserTodos)
		v1.GET("/applied-jobs", middlewares.RequesterIsEmployee(), controllers2.AppliedJobs)

		v1.POST("/:job_id/apply-job", middlewares.RequesterIsEmployee(), controllers2.ApplyToTheJob)

		v1.POST("/create", middlewares.RequesterIsEmployer(), controllers2.CreateJob)
		v1.PUT("/update/:job_id", middlewares.RequesterIsEmployer(), controllers2.UpdateJob)

		v1.GET("/applicants", middlewares.RequesterIsEmployer(), controllers2.ApplicantsForEmployer)
	}
}
