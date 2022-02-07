package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manjurulhoque/golang-job-portal/controllers"
	"github.com/manjurulhoque/golang-job-portal/middlewares"
)

func (r routes) addJobRoutes(rg *gin.RouterGroup) {
	v1 := rg.Group("/jobs")

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

	v2 := v1
	v2.Use(middlewares.AuthMiddleware())
	{
		v2.Use(middlewares.RequesterIsEmployee())
		{
			v2.POST("/:job_id/apply-job", controllers.ApplyToTheJob)
		}

		v2.GET("/user", controllers.CurrentUserTodos)
	}
}
