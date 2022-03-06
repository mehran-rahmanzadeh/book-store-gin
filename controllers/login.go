package controllers

import (
	"gin-tutorial/utils"
	"github.com/gin-gonic/gin"
)

// LoginController login controller interface
type LoginController interface {
	Login(ctx *gin.Context) string
}

// LoginCredentials Login credential
type LoginCredentials struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type LoginControllerService struct {
	loginService utils.LoginService
	jWtService   utils.JWTService
}

func LoginHandler(loginService utils.LoginService,
	jWtService utils.JWTService) LoginController {
	return &LoginControllerService{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *LoginControllerService) Login(ctx *gin.Context) string {
	var credential LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isUserAuthenticated := controller.loginService.LoginUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		return controller.jWtService.GenerateToken(credential.Email, true)

	}
	return ""
}
