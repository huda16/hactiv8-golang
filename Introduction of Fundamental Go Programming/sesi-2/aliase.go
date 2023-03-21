package main

import "fmt"

func main() {
	// declare the variables with uint8 data type
	var a uint8 = 10
	var b byte // byte is an aliase of uint8 data type

	b = a // no error, because byte has the same data type with uint8
	_ = b

	// declare aliase data type named second
	// type <aliase_name> = <data_type_name>
	type second = uint

	var hour second = 3600
	fmt.Printf("hour type: %T\n", hour)
}
