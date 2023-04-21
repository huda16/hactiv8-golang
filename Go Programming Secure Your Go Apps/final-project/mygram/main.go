package main

import (
	"mygram/database"
	"mygram/router"
)

func main() {
	var PORT = ":8080"

	database.StartDB()
	router.StartServer().Run(PORT)
}
