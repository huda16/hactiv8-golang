package main

import (
	"challenge-2/database"
	"challenge-2/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
