package main

import "fmt"

func main() {
	// the first way
	// for i := 0; i < 3; i++ {
	// 	fmt.Println("Angka", i)
	// }

	// the second way
	// var i = 0

	// for i < 3 {
	// 	fmt.Println("Angka", i)
	// 	i++
	// }

	// the third way
	// var i = 0

	// for {
	// 	fmt.Println("Angka", i)

	// 	i++
	// 	if i == 3 {
	// 		break
	// 	}
	// }

	// break and continue keyword
	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 1 {
	// 		continue
	// 	}

	// 	if i > 8 {
	// 		break
	// 	}

	// 	fmt.Println("Angka", i)
	// }

	// nested looping
	// for i := 0; i < 5; i++ {
	// 	for j := i; j < 5; j++ {
	// 		fmt.Print(j, " ")
	// 	}

	// 	fmt.Println()
	// }

	// labeling
outerLoop:
	for i := 0; i < 3; i++ {
		fmt.Println("Perulangan ke -", i+1)
		for j := 0; j < 3; j++ {
			if i == 2 {
				break outerLoop
			}
			fmt.Print(j, " ")
		}
		fmt.Print("\n")
	}
}
