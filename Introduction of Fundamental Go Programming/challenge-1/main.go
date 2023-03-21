package main

import "fmt"

func main() {
	i := 21
	j := true
	ya := 'Я'
	k := 123.456

	// menampilkan nilai i
	fmt.Printf("%v\n", i)

	// menampilkan tipe data dari variabel i
	fmt.Printf("%T\n", i)

	// menampilkan tanda %
	fmt.Printf("%%\n")

	// menampilkan nilai boolean j
	fmt.Printf("%v\n", j)
	
	// menampilkan unicode russia Я (ya)
	fmt.Printf("%v\n", ya)

	// menampilkan nilai base 10
	fmt.Printf("%d\n", i)

	// menampilkan nilai base 8
	fmt.Printf("%o\n", i)

	// menampilkan nilai base 16
	hexValue := fmt.Sprintf("%x", ya)
	hexValue = hexValue[2:]
	fmt.Println(hexValue)

	// menampilkan nilai base 16
	hexValueF := fmt.Sprintf("%X", ya)
	hexValueF = hexValueF[2:]
	fmt.Println(hexValueF)

	// menampilkan unicode karakter Я
	fmt.Printf("%U\n", ya)

	// menampilkan float
	fmt.Printf("%f\n", k)

	// Menampilkan float scientific
	fmt.Printf("%e\n", k)
}
