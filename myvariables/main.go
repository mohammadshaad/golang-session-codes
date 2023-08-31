package main

import "fmt"

func main() {
	// variables declaration

	var username1 string = "mohammadshaad"
	fmt.Println("Username1 is ", username1)

	var username2 string
	username2 = "xyz"
	fmt.Println("Username2 is ", username2)

	// const keyword
	const password = "shaad"
	fmt.Println("Password ", password)

	var phoneNumber int = 123456
	fmt.Println("Phone Number : ", phoneNumber)

	var uintExample uint = 123455.0
	fmt.Println("This is my unint number ", uintExample)

	var verified bool = false
	fmt.Println(verified)

	fmt.Printf("The type of verified variable is : %T\n", password)

	var smallVal float64 = 400.2323492374982374982374892374832894
	fmt.Println(smallVal)


	// Most commonly most widely and easiest way of declaring variables
	username := true
	fmt.Println("My username is : ", username)
	fmt.Printf("The type of username variable is : %T\n" , username)


}
