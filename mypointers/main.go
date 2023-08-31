package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers!")

	// oldNumbers := 40

	// fmt.Println("Old Number is : ", oldNumbers)

	// newNumber := &oldNumbers // & - Reference Operator
	// fmt.Println("New Number : ", *newNumber) //delimiter
	// fmt.Printf("The type of newNumber is : %T\n", newNumber)

	// fmt.Println(newNumber)

	// myNumber := 100

	// ptr := &myNumber

	// fmt.Println("My number is : ", myNumber)

	// *ptr = *ptr + 1

	// fmt.Println("The value of the new number is : ", *ptr)
	// fmt.Println("The value of the new number is : ", myNumber)
	// fmt.Printf("The type of newNumber is : %T\n", ptr)


	anotherVariable := 50

	ptr := &anotherVariable

	fmt.Println("The address of anotherVariable is : ", ptr)
	fmt.Println("The value of anotherVariable is : ", *ptr)

	*ptr = *ptr * 100
	fmt.Println("The modifed value is :  ", anotherVariable)


}
