package models

import "gorm.io/gorm"

// Book model info
// @Description Book information
type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
}

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type PDF struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Size        uint   `json:"size"`
}

type CreatePDFInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Size        uint   `json:"size" binding:"required"`
}

type UpdatePDFInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Size        uint   `json:"size"`
}
