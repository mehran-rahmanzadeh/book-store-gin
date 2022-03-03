package main

import (
	"gin-tutorial/controllers"
	"gin-tutorial/models"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// initiate default router
	router = gin.Default()

	// connect to db
	models.ConnectDatabase()

	// routes
	router.GET("/books/", controllers.BooksList)
	router.GET("/books/:id/", controllers.BookDetail)
	router.POST("/books/", controllers.CreateBook)
	router.PATCH("/books/:id/", controllers.UpdateBook)
	router.DELETE("/books/:id/", controllers.DeleteBook)

	router.GET("/pdfs/", controllers.PDFList)
	router.GET("/pdfs/:id/", controllers.PDFDetail)
	router.POST("/pdfs/", controllers.CreatePDF)
	router.PATCH("/pdfs/:id/", controllers.UpdatePDF)
	router.DELETE("/pdfs/:id/", controllers.DeletePDF)

	// start serving the application
	err := router.Run()
	if err != nil {
		return
	}
}
