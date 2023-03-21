package main

import "fmt"

func main() {
	// int dan uint
	// var first = 89
	// var second = -12

	var first uint8 = 89
	var second int8 = -12

	fmt.Printf("Tipe data first : %T\n", first)
	fmt.Printf("Bilangan second : %T\n", second)

	// float
	var decimalNumber float32 = 3.63

	fmt.Printf("decimal Number: %f\n", decimalNumber)
	fmt.Printf("decimal Number: %.3f\n", decimalNumber)

	// bool
	var condition bool = true
	fmt.Printf("is it permitted? %t \n", condition)

	// string
	var message string = "Halo"
	message = `Selamat datang di "Hactive8".
	Salam kenal :).
	Mari belajar "Scalable Web Service With Go".`
	fmt.Printf(message)
	fmt.Println(message)
}
