package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/manjurulhoque/golang-job-portal/handlers"
	"github.com/manjurulhoque/golang-job-portal/models"
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
// @Router /users/profile [get]
func UserProfile(c *gin.Context) {
	user, _ := utils.AuthorizedUser(c)

	c.JSON(http.StatusOK, utils.SuccessResponse(user))
}

// UpdateUserProfile godoc
// @Summary Update User profile.
// @Description Update User profile.
// @Tags user, profile
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users/profile [post]
func UpdateUserProfile(c *gin.Context) {
	var userModel models.UpdateUserProfile

	if err := c.ShouldBindJSON(&userModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid_data", "message": err.Error()})
		c.Abort()
		return
	}

	errs := utils.TranslateError(userModel)
	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs})
		return
	}

	user, _ := utils.AuthorizedUser(c)

	err := handlers.UpdateUserProfile(c, &userModel, &user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Successfully updated user profile"})
	}
}
