package main

import (
	"fmt"
	"strings"
)

func main() {
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}

	var words = [3]string{"Airell", "Nanda", "Mailo"}

	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", words)

	// modify element through index
	var fruits = [3]string{"apel", "pisang", "mangga"}
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "mango"

	fmt.Printf("%#v\n", fruits)

	// loop through element
	// the first way
	for i, v := range fruits {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}

	fmt.Println(strings.Repeat("#", 25))

	// the second way
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Index: %d, Value: %s\n", i, fruits[i])
	}

	// multidimensional array
	balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}
	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
}
