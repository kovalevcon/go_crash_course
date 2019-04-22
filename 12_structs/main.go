package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	//firstName string
	//lastName string
	//city string
	//gender string
	//age int

	// Short way
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// HasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Constantine", lastName: "Kovalev", city: "Almaty", gender: "m", age: 29}
	// Alternative
	person2 := Person{"Anna", "Bassova", "Almaty", "f", 30}

	fmt.Println(person1, person2)
	fmt.Println(person1.firstName)

	person1.age++
	person1.hasBirthday()

	fmt.Println(person1.greet())

	person2.getMarried(person1.lastName)
	fmt.Println(person2)
}
