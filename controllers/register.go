package controllers

import (
	"errors"
	"gin-tutorial/models"
	"gin-tutorial/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// Register godoc
// @Summary      Register new user
// @Description  signup user
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.UserRegisterInput
// @Failure      400  {object}	utils.ErrorMessage
// @Failure      500  {object}	utils.ErrorMessage
// @Router       /register/ [post]
func Register(c *gin.Context) {
	var input models.UserRegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			errorMessages := make([]utils.ErrorMessage, len(validationErrors))
			for idx, field := range validationErrors {
				errorMessages[idx] = utils.ErrorMessage{
					Field:   field.Field(),
					Message: utils.GetErrorMessage(field),
				}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		}
		return
	}

	var user models.User = models.User{Email: input.Email, Password: input.Password}
	models.DB.Create(&user)

	c.JSON(http.StatusCreated, gin.H{"data": "user registered successfully"})
}
