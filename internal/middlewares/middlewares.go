package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/manjurulhoque/golang-job-portal/internal/constants"
	"github.com/manjurulhoque/golang-job-portal/internal/controllers"
	"github.com/manjurulhoque/golang-job-portal/internal/handlers"
	"github.com/manjurulhoque/golang-job-portal/pkg/utils"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		bearerToken := ""

		if len(strings.Split(token, " ")) == 2 {
			bearerToken = strings.Split(token, " ")[1]
		}

		if bearerToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}

		claims, err := controllers.VerifyAction(bearerToken)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		user, exists := handlers.FindUserByEmail(claims.Email)
		if !exists || user.Email != claims.Email {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Set("Claims", claims)
		c.Set("AuthorizedUser", user)
		c.Next()
	}
}

// RequesterIsEmployee .
func RequesterIsEmployee() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := utils.AuthorizedUser(c)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		if user.Role != constants.EmployeeRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You are not authorized to access this resource"})
			return
		}

		c.Set("RequesterIsAuthorizedUser", true)
		c.Next()
	}
}

// RequesterIsEmployer .
func RequesterIsEmployer() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := utils.AuthorizedUser(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		if user.Role != constants.EmployerRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You are not authorized to access this resource"})
			return
		}

		c.Set("RequesterIsAuthorizedUser", true)
		c.Next()
	}
}
