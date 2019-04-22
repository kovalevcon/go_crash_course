package main

import "fmt"

func main() {
	// Arrays
	//var fruitArr [2]string
	//fruitArr[0] = "Apple"
	//fruitArr[1] = "Orange"

	//fruitArr := [2]string{"Apple", "Orange"}
	//
	//fmt.Println(fruitArr[0])

	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}
