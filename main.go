package main

import (
	"gin-tutorial/controllers"
	"gin-tutorial/docs"
	"gin-tutorial/middlewares"
	"gin-tutorial/models"
	"gin-tutorial/utils"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

var router *gin.Engine

// @title        Go + Gin Books API
// @version      1.0
// @description  This is a sample server books server. You can visit the GitHub repository at https://github.com/mehran-rahmanzadeh/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host                     localhost:8080
// @BasePath                 /
// @query.collection.format  multi
// @securityDefinitions.apikey ApiKeyAuth
// @tokenUrl http://localhost:8080/login/
// @in header
// @name Authorization

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
	secureRouter.PATCH("/profile/edit/", controllers.EditProfile)

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

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// start serving the application
	err := router.Run()
	if err != nil {
		return
	}
}
