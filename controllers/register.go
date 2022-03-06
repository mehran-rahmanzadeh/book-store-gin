package controllers

import (
	"gin-tutorial/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var input models.UserRegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User = models.User{Email: input.Email, Password: input.Password}
	models.DB.Create(&user)

	c.JSON(http.StatusCreated, gin.H{"data": "user registered successfully"})
}
