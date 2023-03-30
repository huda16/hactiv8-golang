package main

import "fmt"

func main() {
	// initiate variable n
	var n int
	fmt.Print("Masukkan nilai n: ")
	// take the value of variable n from input user
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	} 
}