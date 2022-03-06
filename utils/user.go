package utils

import (
	"gin-tutorial/models"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) (models.User, bool) {
	var user models.User
	email, isExists := c.Get("email")
	if isExists {
		if err := models.DB.Where("email = ?", email).First(&user).Error; err != nil {
			return models.User{}, false
		}
		return user, true
	} else {
		return models.User{}, false
	}
}
