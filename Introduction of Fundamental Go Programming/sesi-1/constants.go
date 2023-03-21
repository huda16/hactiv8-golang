package main

import "fmt"

func main() {
	// constant atau konstanta
	const full_name string = "Mochamad Nurul Huda"

	fmt.Printf("Hello %s\n", full_name)

	// IOTA is number generator for constants which starts from zero, and is incremented by 1 automatically.
	// can be operate with symbols and mathematic operators
	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)

	const (
		c11 = (iota + 2) / 2
		c12 = iota
		c13 = iota
	)

	fmt.Println(c1, c2, c3)

	// operators
	// arithmetic operators +, -, *, /, %, ++, --
	var value = (2 + 2) * 3
	fmt.Println(value)

	// relational operators ==, !=, >, <, >=, <=
	var firstCondition bool = 2 < 3
	var secondCondition bool = "joey" == "Joey"
	var thirdCondition bool = 10 != 2.3
	var fourthCondition bool = 11 <= 11

	fmt.Println("first condition", firstCondition)
	fmt.Println("second condition", secondCondition)
	fmt.Println("third condition", thirdCondition)
	fmt.Println("fourth condition", fourthCondition)

	// logical operators &&, ||, !
	var right = true
	var wrong = false

	var wrongAndRight = wrong && right
	fmt.Printf("wrong && right \t(%t) \n", wrongAndRight)

	var wrongOrRight = wrong || right
	fmt.Printf("wrong || right \t(%t) \n", wrongOrRight)

	var wrongReverse = !wrong
	fmt.Printf("!wrong \t\t(%t) \n", wrongReverse)
}
