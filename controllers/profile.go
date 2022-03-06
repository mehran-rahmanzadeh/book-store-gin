package controllers

import (
	"gin-tutorial/models"
	"gin-tutorial/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Profile(c *gin.Context) {
	user, success := utils.GetUser(c)
	if success {
		c.JSON(http.StatusOK, gin.H{"user": user})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not found."})
	}
}

// TODO: fix there is an error

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
