package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Genre       string `json:"genre"`
	Description string `json:"description"`
	ISBN        string `json:"isbn"`
	Image       string `json:"image"`
	Published   string `json:"published"`
	Publisher   string `json:"publisher"`
}

var books = map[int]Book{}

func GetBooks(c *gin.Context) {
	var result []Book
	for _, book := range books {
		result = append(result, book)
	}
	c.JSON(http.StatusOK, result)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	bookID := 0
	fmt.Sscanf(id, "%d", &bookID)
	book, exists := books[bookID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var newBook Book
	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	if newBook.ID == 0 {
		newBook.ID = len(books) + 1
	}
	books[newBook.ID] = newBook
	c.JSON(http.StatusCreated, gin.H{"id": newBook.ID})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	bookID := 0
	fmt.Sscanf(id, "%d", &bookID)
	var updatedBook Book
	if err := c.BindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}
	_, exists := books[bookID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	updatedBook.ID = bookID
	books[bookID] = updatedBook
	c.JSON(http.StatusOK, gin.H{"message": "Book updated"})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	bookID := 0
	fmt.Sscanf(id, "%d", &bookID)
	_, exists := books[bookID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	delete(books, bookID)
	c.JSON(http.StatusNoContent, nil)
}
