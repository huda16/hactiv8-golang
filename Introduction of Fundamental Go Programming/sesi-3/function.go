package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	greet("Huda", "Jalan Laswi", 21)

	// return
	var names = []string{"Huda", "Satrio"}
	var printMsg = greetReturn("Heii", names)

	fmt.Println(printMsg)

	// returning multiple values
	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Println("Area:", area)
	fmt.Println("Circumference", circumference)

	// variadic function
	studentLists := print("Huda", "Satrio", "Ardi", "Ujang")
	fmt.Printf("%v\n", studentLists)

	// variadic function 2
	numberLists := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := sum(numberLists...)

	fmt.Println("Result:", result)

	// variadic function 3
	profile("Huda", "pasta", "ayam geprek", "ikan roa", "sate padang")
}

func greet(name, address string, age int8) {
	fmt.Printf("Hello there: My name is %s and I'm %d years old", name, age)
	fmt.Println()
}

func greetReturn(msg string, names []string) string {
	var joinStr = strings.Join(names, " ")

	var result = fmt.Sprintf("%s %s", msg, joinStr)

	return result
}

func calculate(d float64) (float64, float64) {
	// menghitung luas
	var area float64 = math.Pi * math.Pow(d/2, 2)

	// menghitung keliling
	var circumference = math.Pi * d

	return area, circumference
}

// predefined return values
// func calculate(d float64) (area float64, circumference float64) {
// 	// menghitung luas
// 	var area float64 = math.Pi * math.Pow(d/2, 2)

// 	// menghitung keliling
// 	var circumference = math.Pi * d

// 	return
// }

// varadic function
func print(names ...string) []map[string]string {
	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("student%d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}

	return result
}

func sum(numbers ...int) int {
	total := 0

	for _, v := range numbers {
		total += v
	}

	return total
}

func profile(name string, favFoods ...string) {
	mergeFavFoods := strings.Join(favFoods, ",")

	fmt.Println("Hello there!!! I'm", name)
	fmt.Println("I really love to eat", mergeFavFoods)
}
