package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to user input file")

	input := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your input : ")

	// Comma OK syntax
	result, err := input.ReadString('\n')

	fmt.Println("Hello Mr. ", result, "This is my output")

	if err != nil {
		panic(err)
	}
}
