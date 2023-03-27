package main

import (
	"fmt"
	"strings"
)

type Employee struct {
	name     string
	age      int
	division string
}

type Person struct {
	name string
	age  int
}

type Worker struct {
	division string
	person   Person
}

func main() {
	var employee Employee

	employee.name = "Huda"
	employee.age = 21
	employee.division = "Software Engineer"

	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.division)

	var employee1 = Employee{}
	employee1.name = "Airell"
	employee1.age = 23
	employee1.division = "Developer"

	var employee2 = Employee{name: "Ananda", age: 23, division: "Finance"}
	fmt.Printf("Employee1: %+v\n", employee1)
	fmt.Printf("Employee2: %+v\n", employee2)

	// pointer to struct
	var employee3 *Employee = &employee2

	fmt.Println("Employee2 name:", employee2.name)
	fmt.Println("Employee3 name:", employee3.name)

	fmt.Println(strings.Repeat("#", 20))

	employee3.name = "Ardi"

	fmt.Println("Employee2 name:", employee2.name)
	fmt.Println("Employee3 name:", employee3.name)

	var worker = Worker{}

	worker.person.name = "Huda"
	worker.person.age = 21
	worker.division = "Curriculum Developer"

	fmt.Printf("%+v", worker)

	// anonymous struct
	var employee4 = struct {
		person   Person
		division string
	}{}

	employee4.person = Person{name: "Huda", age: 21}
	employee4.division = "Curriculum Developer"

	// anonymous struct dengan pengisian field
	var employee5 = struct {
		person   Person
		division string
	}{
		person:   Person{name: "Ananda", age: 23},
		division: "Finance",
	}

	fmt.Printf("Employee4: %+v\n", employee4)
	fmt.Printf("Employee5: %+v\n", employee5)

	// slice to struct
	var people = []Person{
		{name: "Huda", age: 21},
		{name: "Ardi", age: 22},
		{name: "Mailo", age: 23},
	}

	for _, v := range people {
		fmt.Printf("%+v\n", v)
	}

	// slice of anonymous struct
	var employee6 = []struct {
		person   Person
		division string
	}{
		{person: Person{name: "Huda", age: 21}, division: "Software Engineer"},
		{person: Person{name: "Ardi", age: 22}, division: "Software Engineer"},
		{person: Person{name: "Ananda", age: 23}, division: "Finance"},
	}

	for _, v := range employee6 {
		fmt.Printf("%+v\n", v)
	}
}
