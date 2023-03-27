package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) Introduce(msg string) string {
	return fmt.Sprintf("%s My name is %s and I'm %d years old", msg, p.name, p.age)
}

func (p Person) ChangeName1() {
	p.name = "Mailo"
}

func (p *Person) ChangeName2() {
	p.name = "Mailo"
}

func main() {
	var person = Person{name: "Huda", age: 21}

	fmt.Println(person.Introduce("Hello everyone!!"))

	person.ChangeName1()
	fmt.Println("Change name with ChangeName1 method", person.name)

	person.ChangeName2()
	fmt.Println("Change name with ChangeName2 method", person.name)
}
