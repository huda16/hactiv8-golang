package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"project/database"
	"project/models"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

func GetBooks(ctx *gin.Context) {
	db := database.GetDB()
	books := []models.Book{}

	err := db.Find(&books).Error

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, books)
}

func CreateBook(ctx *gin.Context) {
	var newBook models.Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Bad Request",
			"error_message": err.Error(),
		})
		return
	}

	db := database.GetDB()

	Book := models.Book{
		NameBook: newBook.NameBook,
		Author:   newBook.Author,
	}

	err := db.Create(&Book).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Bad Request",
			"error_message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, Book)
}

func UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedBook models.Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	db := database.GetDB()
	book := models.Book{}
	result := db.Model(&book).Where("id = ?", id).Updates(models.Book{NameBook: updatedBook.NameBook, Author: updatedBook.Author})
	err := result.Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Bad Request",
			"error_message": err.Error(),
		})
		return
	}

	count := result.RowsAffected

	if count < 1 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", id),
		})
		return
	}

	number, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Bad Request",
			"error_message": "id must be integer",
		})
		return
	}

	book.ID = uint(number)

	ctx.JSON(http.StatusOK, book)
}

func GetBook(ctx *gin.Context) {
	id := ctx.Param("id")
	db := database.GetDB()

	user := models.Book{}

	result := db.First(&user, "id = ?", id)
	err := result.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error_status":  "Data Not Found",
				"error_message": fmt.Sprintf("book with id %v not found", id),
			})
			return
		}
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")

	db := database.GetDB()
	book := models.Book{}
	result := db.Where("id = ?", id).Delete(&book)
	err := result.Error
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	count := result.RowsAffected

	if count < 1 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Book deleted successfully",
	})
}
