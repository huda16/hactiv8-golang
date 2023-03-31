package main

import "challenge/routers"

func main() {
	var PORT = ":4000"

	routers.StartServer().Run(PORT)
}