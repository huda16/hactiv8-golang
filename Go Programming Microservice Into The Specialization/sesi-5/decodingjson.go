package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func main() {
	decodingJSON()
	decodingArrayOfJSON()
	encodingObjectToJSON()
}

func decodingJSON() {
	var jsonString = `
	{
		"full_name": "Mochamad Nurul Huda",
		"email": "mochamadnurulhuda16@gmail.com",
		"age": 21
	}
	`

	// to struct
	var result Employee
	// to map
	// var result map[string]interface{}

	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// to empty interface
	// var temp interface{}
	// var err = json.Unmarshal([]byte(jsonString), &temp)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// var result = temp.(map[string]interface{})

	fmt.Println("full_name:", result.FullName)
	fmt.Println("email:", result.Email)
	fmt.Println("age:", result.Age)
}

func decodingArrayOfJSON() {
	// to slice of struct
	var jsonString = `[
		{
			"full_name": "Mochamad Nurul Huda",
			"email": "mochamadnurulhuda16@gmail.com",
			"age": 21
		},
		{
			"full_name": "Ardi Hilal",
			"email": "ardihilal@gmail.com",
			"age": 23
		}
	]
	`

	var result []Employee

	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, v := range result {
		fmt.Printf("index %d: %+v\n", i+1, v)
	}
}

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func encodingObjectToJSON() {
	var object = []User{{"John Wick", 27}, {"Ethan Hunt", 32}}

	var jsonData, err = json.Marshal(object)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString = string(jsonData)
	fmt.Println(jsonString)
}
