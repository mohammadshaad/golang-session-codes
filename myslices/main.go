package main

import "fmt"

func main() {
	fmt.Println("Welcome to my slices")

	fruits := []string{"Mango", "Apple", "Pineapple", "Banana"}

	fmt.Println(fruits)

	myNumber := []int{1, 2, 4, 5, 6, 7}

	fmt.Println(myNumber)

	places := make([]string, 4)
	places[0] = "Nashik"
	places[1] = "Chennai"
	places[2] = "Mumbai"

	fmt.Println(places)

	places = append(places, "Kolkata")
	fmt.Println(places)

	// places = append(places[1:])
	// fmt.Println(places)

	places = append(places[:1], places[2:]...)
	fmt.Println(places)

	myNumber = append(myNumber[:1], myNumber[2:]...)
	fmt.Println(myNumber)




}
