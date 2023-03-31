package controllers

import (
	"fmt"
	"strconv"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	BookID int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Desc string `json:"desc"`
}

var BookDatas = []Book{
	{BookID: 1, Title: "Golang", Author: "Gopher", Desc: "A book for Go"},
}

func GetBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, BookDatas)
}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newBook.BookID = len(BookDatas)+1
	BookDatas = append(BookDatas, newBook)

	ctx.JSON(http.StatusCreated, "Created")
}

func UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	condition := false
	var updatedBook Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range BookDatas {
		if id == strconv.Itoa(book.BookID) {
			condition = true
			num, _ := strconv.Atoi(id)
			BookDatas[i] = updatedBook
			BookDatas[i].BookID = num
			break
		}
	}

	if ! condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, "Updated")
}

func GetBook(ctx *gin.Context) {
	id := ctx.Param("id")
	condition := false
	var bookData Book

	for i, book := range BookDatas {
		if id == strconv.Itoa(book.BookID) {
			condition = true
			bookData = BookDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", id),
		})
		return 
	}

	ctx.JSON(http.StatusOK, bookData)
}

func DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	condition := false
	var bookIndex int

	for i, book := range BookDatas {
		if id == strconv.Itoa(book.BookID) {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", id),
		})
		return
	}

	copy(BookDatas[bookIndex:], BookDatas[bookIndex+1:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	ctx.JSON(http.StatusOK, "Deleted")
}