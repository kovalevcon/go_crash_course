package main

import "fmt"

func main() {
	//var name = "Constantine"
	var age = 29
	const isCool = true
	name, email := "Constantine", "kovalevcon@mail.ru"

	fmt.Println(name, email, age, isCool)
	fmt.Printf("%T\n", age)
}
