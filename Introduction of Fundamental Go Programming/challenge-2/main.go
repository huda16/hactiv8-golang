package main

import "fmt"

func main() {
	word := "САШАРВО"

	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i =", i)
	}

	for j := 0; j < 11; j++ {
		if j < 5 {
			fmt.Println("Nilai j =", j)
		}

		if j == 5 {
			for j, char := range word {
				fmt.Printf("character %U '%c' starts at byte position %d\n", char, char, j)
			}
		}

		if j > 5 {
			fmt.Println("Nilai j =", j)
		}
	}
}