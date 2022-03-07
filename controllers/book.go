package controllers

import (
	"gin-tutorial/models"
	"gin-tutorial/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BooksList godoc
// @Summary      Show Books List
// @Security 	 ApiKeyAuth
// @Description  get books list
// @Tags         books
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Book
// @Failure      400  {object}	utils.ErrorMessage
// @Failure      500  {object}	utils.ErrorMessage
// @Router       /books/ [get]
func BooksList(c *gin.Context) {
	var books []models.Book
	allowedFilters := []string{"author", "title"}
	filterMap := map[string]string{}
	for k, v := range c.Request.URL.Query() {
		if utils.CheckItemInSlice(allowedFilters, k) {
			filterMap[k] = v[0]
		}
	}
	models.DB.Where(filterMap).Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// BookDetail godoc
// @Summary      Show Book Detail
// @Security 	 ApiKeyAuth
// @Description  get book detail
// @Tags         books
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Book
// @Failure      400  {object}	utils.ErrorMessage
// @Failure      404  {object}	utils.ErrorMessage
// @Failure      500  {object}	utils.ErrorMessage
// @Router       /books/:id/ [get]
func BookDetail(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// CreateBook godoc
// @Summary      Create Book
// @Security 	 ApiKeyAuth
// @Description  create new book
// @Tags         books
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.CreateBookInput
// @Failure      400  {object}	utils.ErrorMessage
// @Failure      500  {object}	utils.ErrorMessage
// @Router       /books/ [post]
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

// UpdateBook godoc
// @Summary      Update Book Data
// @Security 	 ApiKeyAuth
// @Description  edit book
// @Tags         books
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.UpdateBookInput
// @Failure      400  {object}	utils.ErrorMessage
// @Failure      500  {object}	utils.ErrorMessage
// @Router       /books/:id/ [patch]
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

// DeleteBook godoc
// @Summary      Remove Book
// @Security 	 ApiKeyAuth
// @Description  delete book
// @Tags         books
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Failure      400  {object}	utils.ErrorMessage
// @Failure      500  {object}	utils.ErrorMessage
// @Router       /books/:id/ [delete]
func DeleteBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusAccepted, gin.H{"data": "Deleted successfully."})
}

// PDFList godoc
// @Summary      Show PDFs List
// @Security 	 ApiKeyAuth
// @Description  get pdfs list
// @Tags         pdfs
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.PDF
// @Failure      400  {object}	utils.ErrorMessage
// @Failure      500  {object}	utils.ErrorMessage
// @Router       /pdfs/ [get]
func PDFList(c *gin.Context) {
	var pdfs []models.PDF
	allowedFilters := []string{"size", "title"}
	filterMap := map[string]string{}
	for k, v := range c.Request.URL.Query() {
		if utils.CheckItemInSlice(allowedFilters, k) {
			filterMap[k] = v[0]
		}
	}
	models.DB.Where(filterMap).Find(&pdfs)

	c.JSON(http.StatusOK, gin.H{"data": pdfs})
}

// PDFDetail godoc
// @Summary      Show PDF Detail
// @Security 	 ApiKeyAuth
// @Description  get pdf detail
// @Tags         pdfs
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.PDF
// @Failure      400  {object}	utils.ErrorMessage
// @Failure      500  {object}	utils.ErrorMessage
// @Router       /pdfs/:id/ [get]
func PDFDetail(c *gin.Context) {
	var pdf models.PDF
	if err := models.DB.Where("id = ?", c.Param("id")).First(&pdf).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pdf})
}

// CreatePDF godoc
// @Summary      Create PDF
// @Security 	 ApiKeyAuth
// @Description  create new pdf
// @Tags         pdfs
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.CreatePDFInput
// @Failure      400  {object}	utils.ErrorMessage
// @Failure      500  {object}	utils.ErrorMessage
// @Router       /pdfs/ [post]
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

// UpdatePDF godoc
// @Summary      Update PDF Data
// @Security 	 ApiKeyAuth
// @Description  edit pdf
// @Tags         pdfs
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.UpdatePDFInput
// @Failure      400  {object}	utils.ErrorMessage
// @Failure      500  {object}	utils.ErrorMessage
// @Router       /pdfs/:id/ [patch]
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

// DeletePDF godoc
// @Summary      Delete PDF
// @Security 	 ApiKeyAuth
// @Description  remove pdf
// @Tags         pdfs
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Failure      400  {object}	utils.ErrorMessage
// @Failure      500  {object}	utils.ErrorMessage
// @Router       /pdfs/:id/ [delete]
func DeletePDF(c *gin.Context) {
	var pdf models.PDF
	if err := models.DB.Where("id = ?", c.Param("id")).First(&pdf).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&pdf)

	c.JSON(http.StatusAccepted, gin.H{"data": "Deleted Successfully."})
}
