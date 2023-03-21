package main

import (
	"fmt"
	"strings"
)

func main() {
	// the first way
	// var fruits = []string{"apple", "banana", "mango"}

	// the second way (make function)
	var fruits = make([]string, 3)
	_ = fruits

	// append function
	// the first way
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "mango"

	// the second way
	fruits = append(fruits, "apple", "banana", "mango")

	fmt.Printf("%#v\n", fruits)

	// the third way with ellipsis (...)
	var fruits1 = []string{"apple", "banana", "mango"}
	var fruits2 = []string{"durian", "pineapple", "starfruit"}

	fruits1 = append(fruits1, fruits2...)

	fmt.Printf("%#v\n", fruits1)

	// copy function
	nn := copy(fruits1, fruits2)

	fmt.Println("Fruits1 =>", fruits1)
	fmt.Println("Fruits2 =>", fruits2)
	fmt.Println("Copied elements =>", nn)

	// slicing
	var fruits3 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	var fruits4 = fruits3[1:4]
	fmt.Printf("%#v\n", fruits4)

	var fruits5 = fruits3[0:]
	fmt.Printf("%#v\n", fruits5)

	var fruits6 = fruits3[:3]
	fmt.Printf("%#v\n", fruits6)

	var fruits7 = fruits3[:] // equals to fruits3[:len(fruits3)]
	fmt.Printf("%#v\n", fruits7)

	// combining slicing and append
	fruits3 = append(fruits3[:3], "rambutan")
	fmt.Printf("%#v\n", fruits3)

	// backing array
	var fruits8 = []string{"apple", "banana", "mango", "durian", "starfruit"}
	var fruits9 = fruits8[2:4]
	fruits9[0] = "rambutan"

	fmt.Println("fruits8 =>", fruits8)
	fmt.Println("fruits9 =>", fruits9)

	// cap function
	var fruits10 = []string{"apple", "mango", "durian", "banana"}

	fmt.Println("Fruits10 cap:", cap(fruits10))
	fmt.Println("Fruits10 len:", len(fruits10))

	fmt.Println(strings.Repeat("#", 20))

	var fruits11 = fruits10[0:3]

	fmt.Println("Fruits11 cap:", cap(fruits11))
	fmt.Println("Fruits11 len:", len(fruits11))

	fmt.Println(strings.Repeat("#", 20))

	var fruits12 = fruits10[1:]

	fmt.Println("Fruits12 cap:", cap(fruits12))
	fmt.Println("Fruits12 len:", len(fruits12))

	// creating a new backing array
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"
	fmt.Println("cars:", cars)
	fmt.Println("newCars:", newCars)
}
