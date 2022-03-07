package controllers

import (
	"gin-tutorial/models"
	"gin-tutorial/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Profile godoc
// @Summary      Show user's profile
// @Security 	 ApiKeyAuth
// @Description  get user's profile
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.User
// @Failure      400  {object}	utils.ErrorMessage
// @Failure      500  {object}	utils.ErrorMessage
// @Router       /profile/ [get]
func Profile(c *gin.Context) {
	user, success := utils.GetUser(c)
	if success {
		c.JSON(http.StatusOK, gin.H{"user": user})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not found."})
	}
}

// EditProfile TODO: fix there is an error
// EditProfile godoc
// @Summary      Edit user's profile
// @Security 	 ApiKeyAuth
// @Description  edit user profile
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.UserEditInput
// @Failure      400  {object}	utils.ErrorMessage
// @Failure      500  {object}	utils.ErrorMessage
// @Router       /profile/ [patch]
func EditProfile(c *gin.Context) {
	user, success := utils.GetUser(c)
	if success {
		var input models.UserEditInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		models.DB.Model(&user).Updates(input)
		c.JSON(http.StatusOK, gin.H{"user": user})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not found."})
	}
}
