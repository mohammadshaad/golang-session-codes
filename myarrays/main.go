package main

import "fmt"

func main() {
	fmt.Println("Welcome to my arrays !!!!")

	fruits := [4]string{"Mango", "Apple"}
	fmt.Println(fruits)

	var places [4]string
	places[0] = "Nashik"
	places[1] = "Chennai"
	places[2] = "Mumbai"
	places[3] = "Kolkata"

	fmt.Println(places)

	myNumbers := [2]int{34}
	fmt.Println(myNumbers)
	fmt.Println(myNumbers[1])

}
