package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/manjurulhoque/golang-job-portal/utils"
	"net/http"
)

// UserProfile godoc
// @Summary User profile.
// @Description User profile.
// @Tags user, profile
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users/profile [post]
func UserProfile(c *gin.Context) {
	user, _ := utils.AuthorizedUser(c)

	c.JSON(http.StatusOK, utils.SuccessResponse(user))
}
