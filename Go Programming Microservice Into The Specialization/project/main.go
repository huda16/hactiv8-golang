package main

import (
	"project/database"
	"project/routers"
)

func main() {
	var PORT = ":4000"

	database.StartDB()
	routers.StartServer().Run(PORT)
}
