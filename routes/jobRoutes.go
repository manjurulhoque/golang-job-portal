package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manjurulhoque/golang-job-portal/controllers"
)

func (r routes) addJobRoutes(rg *gin.RouterGroup) {
	v1 := rg.Group("/jobs")

	v1.POST("/create", controllers.CreateJob)
}
