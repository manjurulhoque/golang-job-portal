package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manjurulhoque/golang-job-portal/controllers"
	"github.com/manjurulhoque/golang-job-portal/middlewares"
)

func (r routes) addUserRoutes(rg *gin.RouterGroup) {
	v1 := rg.Group("/users")

	v1.Use(middlewares.AuthMiddleware())
	{
		v1.GET("/profile", controllers.UserProfile)
		v1.POST("/profile", controllers.UpdateUserProfile)
	}
}
