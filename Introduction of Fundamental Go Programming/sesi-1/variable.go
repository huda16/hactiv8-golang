package main

import "fmt"

func main() {
	// var name string = "Airell"
	// var age int = 23

	// menetapkan kembali nilai variabel
	// name = "Huda"
	// age = 21

	// menatapkan kembali nilai variabel dengan cara yang salah
	// name = 30
	// age = "Huda"

	// fmt.Println("Ini adalah namanya ==>", name)
	// fmt.Println("Ini adalah umurnya nya ==>", age)

	// Variabel tanpa tipe data / type interface
	// var name = "Huda"
	// var age = 21

	// Variabel tanpa tipe data (short declaration)
	// name := "Huda"
	// age := 21

	// fmt.Printf("%T, %T", name, age)
	// fmt.Println("\n")

	// multiple declaration variable
	// var student1, student2, student3 string = "satu", "dua", "tiga"
	// var first, second, third int
	// first, second, third = 1, 2, 3

	// fmt.Println(student1, student2, student3)
	// fmt.Println(first, second, third)

	// multiple declaration variable with type interface and short declaration
	// var name, age, address = "Huda", 21, "Jalan Laswi"
	// first, second, third := "1", 2, "3"

	// fmt.Println(name, age, address)
	// fmt.Println(first, second, third)

	// underscore variable (mengatasi error variabel tidak dipakai)
	// var firstVariable string
	// var name, age, address = "Huda", 21, "Jalan Laswi"

	// _, _, _, _ = firstVariable, name, age, address

	// penggunaan fmt.Printf()
	first, second := 1, "2"

	fmt.Printf("Tipe data variable first adalah %T \n", first)
	fmt.Printf("Tipe data variable second adalah %T \n", second)

	var name = "Huda"
	var age = 21
	var address = "Jalan Laswi"

	fmt.Printf("Halo nama ku %s, umur aku adalah %d, dan aku tinggal di %s", name, age, address)

}
