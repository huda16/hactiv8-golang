package main

import "challenge-3/routers"

func main() {
	var PORT = ":4000"

	routers.StartServer().Run(PORT)
}
