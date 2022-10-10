package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manjurulhoque/golang-job-portal/controllers"
	"github.com/manjurulhoque/golang-job-portal/middlewares"
)

func (r routes) addJobRoutes(rg *gin.RouterGroup) {
	v1 := rg.Group("/jobs")

	v1.GET("/", controllers.AllJobs)
	v1.GET("/:job_id", controllers.JobDetails)

	v1.Use(middlewares.AuthMiddleware())
	{
		v1.GET("/user", controllers.CurrentUserTodos)
		v1.GET("/applied-jobs", middlewares.RequesterIsEmployee(), controllers.AppliedJobs)

		v1.POST("/:job_id/apply-job", middlewares.RequesterIsEmployee(), controllers.ApplyToTheJob)

		v1.POST("/create", middlewares.RequesterIsEmployer(), controllers.CreateJob)
		v1.PUT("/update/:job_id", middlewares.RequesterIsEmployer(), controllers.UpdateJob)

		v1.GET("/applicants", middlewares.RequesterIsEmployer(), controllers.ApplicantsForEmployer)
	}
}
