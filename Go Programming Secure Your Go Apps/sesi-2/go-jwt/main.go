package main

import (
	"go-jwt/database"
	"go-jwt/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
