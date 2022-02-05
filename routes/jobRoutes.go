package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manjurulhoque/golang-job-portal/controllers"
	"github.com/manjurulhoque/golang-job-portal/middlewares"
)

func (r routes) addJobRoutes(rg *gin.RouterGroup) {
	v1 := rg.Group("/jobs")

	//v1.POST("/create", controllers.CreateJob)
	v1.Use(middlewares.AuthMiddleware())
	{
		v1.POST("/create", controllers.CreateJob)
		v1.PUT("/update/:job_id", controllers.UpdateJob)
		v1.GET("/user", controllers.CurrentUserTodos)
	}
}
