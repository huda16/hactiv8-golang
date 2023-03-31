package main

import "gin-framework/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}