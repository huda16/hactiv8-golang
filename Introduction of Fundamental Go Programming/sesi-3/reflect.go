package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func (s student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama		:", reflectType.Field(i).Name)
		fmt.Println("tipe data	:", reflectType.Field(i).Type)
		fmt.Println("nilai		:", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("tipe variabel :", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel :", reflectValue.Int())
	}

	// access value with interface
	fmt.Println("nilai variabel :", reflectValue.Interface())

	// var nilai = reflectValue.Interface().(int)

	// access information variables object of property
	var s1 = &student{Name: "Aduh", Grade: 3}
	s1.getPropertyInfo()

	// access information variables object of method
	var s2 = &student{Name: "John Wick", Grade: 2}
	fmt.Println("nama :", s2.Name)

	var reflectValueS2 = reflect.ValueOf(s2)
	var method = reflectValueS2.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("aaaa"),
	})

	fmt.Println("nama :", s2.Name)
}
