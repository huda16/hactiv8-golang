package controllers

import (
	"challenge-3/models"
	"database/sql"
	"fmt"
	"net/http"
	_ "strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"desc"`
}

var (
	db  *sql.DB = models.Connection()
	err error
)

func GetBooks(ctx *gin.Context) {
	var results = []Book{}

	sqlStatement := `SELECT * from books`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var book = Book{}

		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description)

		if err != nil {
			panic(err)
		}

		results = append(results, book)
	}

	ctx.JSON(http.StatusOK, results)
}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	sqlStatement := `INSERT INTO books (title, author, description) VALUES ($1, $2, $3) Returning *`

	err = db.QueryRow(sqlStatement, newBook.Title, newBook.Author, newBook.Description).Scan(&newBook.ID, &newBook.Title, &newBook.Author, &newBook.Description)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, newBook)
}

func UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedBook Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	sqlStatement := `UPDATE books SET title = $2, author = $3, description = $4 WHERE id = $1;`
	res, err := db.Exec(sqlStatement, id, updatedBook.Title, updatedBook.Author, updatedBook.Description)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	count, err := res.RowsAffected()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if count < 1 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %v has been successfully updated", id),
	})
}

func GetBook(ctx *gin.Context) {
	id := ctx.Param("id")
	var newBook Book

	sqlStatement := `SELECT * FROM books WHERE id = $1;`
	err := db.QueryRow(sqlStatement, id).Scan(&newBook.ID, &newBook.Title, &newBook.Author, &newBook.Description)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, newBook)
}

func DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")

	sqlStatement := `DELETE FROM books WHERE id = $1;`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	count, err := res.RowsAffected()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if count < 1 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %v has been successfully deleted", id),
	})
}
