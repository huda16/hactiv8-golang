package main

import (
	"challenge-3/database"
	"challenge-3/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
