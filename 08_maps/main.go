package main

import "fmt"

func main() {
	//// Define map
	//emails := make(map[string]string)
	//
	//// Assign kv
	//emails["Kovalevcon"] = "kovalevcon@mail.ru"
	//emails["Zilich"] = "zilich08@gmail.com"
	//emails["Del"] = "del@temp.com"

	// Declare map and add kv
	emails := map[string]string{"Kovalevcon": "kovalevcon@mail.ru", "Zilich": "zilich08@gmail.com", "Del": "del@temp.com"}

	fmt.Println(emails)
	fmt.Println(emails["Kovalevcon"])
	fmt.Println(len(emails))

	// Delete from map
	delete(emails, "Del")
	fmt.Println(emails)
}
