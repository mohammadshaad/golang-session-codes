package main

import "fmt"

func main() {
	fmt.Println("Welcome to my functions!!!")

	statement, result := addNum(1, 2)
	fmt.Println(statement, result)

}

// function signatures - func nameOfFunction (parameteres) type of function (optional)
func addNum(a int, b int) (string, int) {
	return "The result is ", b + a
}


