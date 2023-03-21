package main

import (
	"fmt"
)

func main() {
	// the first way
	// declaration
	// var person map[string]string

	// initialisation
	// person = map[string]string{}

	// person["name"] = "Huda"
	// person["age"] = "21"
	// person["address"] = "Jalan Laswi"

	// the second way
	var person = map[string]string{
		"name":    "Huda",
		"age":     "21",
		"address": "Jalan Sudirman",
	}

	fmt.Println("name:", person["name"])
	fmt.Println("age:", person["age"])
	fmt.Println("address:", person["address"])

	// looping with map
	for key, value := range person {
		fmt.Println(key, ":", value)
	}

	// deleting value
	// fmt.Println("Before deleting:", person)
	// delete(person, "age")
	// fmt.Println("After deleting:", person)

	value, exist := person["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value doesn't exist")
	}

	delete(person, "age")

	value, exist = person["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value has been deleted")
	}

	var peoples = []map[string]string{
		{"name": "Huda", "age": "21"},
		{"name": "Satrio", "age": "23"},
		{"name": "Ujang", "age": "25"},
	}

	for i, people := range peoples {
		fmt.Printf("Index: %d, name: %s, age: %s\n", i, people["name"], people["age"])
	}
}
