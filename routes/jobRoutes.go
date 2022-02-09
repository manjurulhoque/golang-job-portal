package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manjurulhoque/golang-job-portal/controllers"
	"github.com/manjurulhoque/golang-job-portal/middlewares"
)

func (r routes) addEmployeeJobRoutes(rg *gin.RouterGroup) {
	v1 := rg.Group("/employee/jobs")

	v1.GET("/", controllers.AllJobs)

	v1.Use(middlewares.AuthMiddleware())
	{
		v1.GET("/user", controllers.CurrentUserTodos)
		v1.GET("/applied-jobs", controllers.AppliedJobs)

		v1.Use(middlewares.RequesterIsEmployee())
		{
			v1.POST("/:job_id/apply-job", controllers.ApplyToTheJob)
		}
	}
}

func (r routes) addEmployerJobRoutes(rg *gin.RouterGroup) {
	v1 := rg.Group("/employer/jobs")

	v1.GET("/", controllers.AllJobs)

	//v1.POST("/create", controllers.CreateJob)
	v1.Use(middlewares.AuthMiddleware())
	{
		v1.Use(middlewares.RequesterIsEmployer())
		{
			v1.POST("/create", controllers.CreateJob)
			v1.PUT("/update/:job_id", controllers.UpdateJob)
		}
	}
}
