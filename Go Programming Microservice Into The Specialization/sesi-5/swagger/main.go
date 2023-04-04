package main

import (
	"swagger/database"
	"swagger/routers"
)

func main() {
	var PORT = ":8080"

	database.StartDB()
	routers.StartServer().Run(PORT)
}
