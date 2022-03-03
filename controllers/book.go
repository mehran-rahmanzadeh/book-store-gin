package controllers

import (
	"gin-tutorial/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BooksList(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func BookDetail(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func CreateBook(c *gin.Context) {
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var book models.Book = models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusCreated, gin.H{"data": book})
}

func UpdateBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&book).Updates(input)

	c.JSON(http.StatusAccepted, gin.H{"data": book})
}

func DeleteBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusAccepted, gin.H{"data": "Deleted successfully."})
}

func PDFList(c *gin.Context) {
	var pdfs []models.PDF
	models.DB.Find(&pdfs)

	c.JSON(http.StatusOK, gin.H{"data": pdfs})
}

func PDFDetail(c *gin.Context) {
	var pdf models.PDF
	if err := models.DB.Where("id = ?", c.Param("id")).First(&pdf).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pdf})
}

func CreatePDF(c *gin.Context) {
	var input models.CreatePDFInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var pdf models.PDF = models.PDF{Title: input.Title, Description: input.Description, Size: input.Size}
	models.DB.Create(&pdf)

	c.JSON(http.StatusCreated, gin.H{"data": pdf})
}

func UpdatePDF(c *gin.Context) {
	var pdf models.PDF
	if err := models.DB.Where("id = ?", c.Param("id")).First(&pdf).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	var input models.UpdatePDFInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&pdf).Updates(input)

	c.JSON(http.StatusAccepted, gin.H{"data": pdf})
}

func DeletePDF(c *gin.Context) {
	var pdf models.PDF
	if err := models.DB.Where("id = ?", c.Param("id")).First(&pdf).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&pdf)

	c.JSON(http.StatusAccepted, gin.H{"data": "Deleted Successfully."})
}
