package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manjurulhoque/golang-job-portal/controllers"
	"github.com/manjurulhoque/golang-job-portal/middlewares"
)

func (r routes) addTagRoutes(rg *gin.RouterGroup) {
	v1 := rg.Group("/tags")
	v1.GET("/", controllers.AllTags)

	//v1.POST("/create", controllers.CreateJob)
	v1.Use(middlewares.AuthMiddleware())
	{
		v1.POST("/create", controllers.CreateTag)
	}
}
