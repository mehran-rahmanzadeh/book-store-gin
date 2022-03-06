package controllers

import (
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
