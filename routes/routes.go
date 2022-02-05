package routes

import (
	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

func SetupRouter() *gin.Engine {
	r := routes{
		router: gin.Default(),
	}

	r.router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	v1 := r.router.Group("/v1/api")
	r.addAuthRoutes(v1)
	r.addJobRoutes(v1)
	r.addTagRoutes(v1)

	return r.router
}
