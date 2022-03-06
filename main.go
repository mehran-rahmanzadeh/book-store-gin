package main

import (
	"gin-tutorial/controllers"
	"gin-tutorial/middlewares"
	"gin-tutorial/models"
	"gin-tutorial/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine

func main() {
	// initiate default router
	router = gin.Default()

	// connect to db
	models.ConnectDatabase()

	// JWT middleware
	secureRouter := router.Group("")
	secureRouter.Use(middlewares.AuthorizeJWT())

	// routes
	secureRouter.GET("/books/", controllers.BooksList)
	secureRouter.GET("/books/:id/", controllers.BookDetail)
	secureRouter.POST("/books/", controllers.CreateBook)
	secureRouter.PATCH("/books/:id/", controllers.UpdateBook)
	secureRouter.DELETE("/books/:id/", controllers.DeleteBook)

	secureRouter.GET("/pdfs/", controllers.PDFList)
	secureRouter.GET("/pdfs/:id/", controllers.PDFDetail)
	secureRouter.POST("/pdfs/", controllers.CreatePDF)
	secureRouter.PATCH("/pdfs/:id/", controllers.UpdatePDF)
	secureRouter.DELETE("/pdfs/:id/", controllers.DeletePDF)

	// profile
	secureRouter.GET("/profile/", controllers.Profile)

	// authentication
	var loginService utils.LoginService = utils.StaticLoginService()
	var jwtService utils.JWTService = utils.JWTAuthService()
	var loginController controllers.LoginController = controllers.LoginHandler(loginService, jwtService)

	router.POST("/login/", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})
	router.POST("/register/", controllers.Register)

	// start serving the application
	err := router.Run()
	if err != nil {
		return
	}
}
