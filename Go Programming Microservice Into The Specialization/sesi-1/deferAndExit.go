package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("defer function starts to execute")
	fmt.Println("Hai everyone")
	fmt.Println("Welcome back to Go Learning Center")
	fmt.Println()
	
	callDeferFunc()
	fmt.Println("Hai everyone!!")
	fmt.Println()

	// exit
	defer fmt.Println("Invoke with defer")
	fmt.Println("Before Exiting")
	os.Exit(1)
}

func callDeferFunc() {
	defer deferFunc()
}

func deferFunc() {
	fmt.Println("defer function starts to execute")
}