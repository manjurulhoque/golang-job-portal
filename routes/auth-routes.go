package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manjurulhoque/golang-job-portal/controllers"
)

func (r routes) addAuthRoutes(rg *gin.RouterGroup) {
	v1 := rg.Group("/users")

	v1.POST("/login", controllers.Login)
	v1.POST("/register", controllers.Register)
}
