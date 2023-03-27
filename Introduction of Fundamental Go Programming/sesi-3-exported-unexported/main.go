package main

import "sesi-3-exported-unexported/helpers"

func main() {
	helpers.Greet()

	// error
	// helpers.greet()

	var person = helpers.Person{}

	person.Invokegreet()
}