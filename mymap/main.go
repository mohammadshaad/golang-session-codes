package main

import "fmt"

func main() {
	fmt.Println("Welcome to my maps!!")

	languages := make(map[string]string)

	languages["Java"] = "Java"
	languages["C++"] = "C++"
	languages["Python"] = "Python"
	languages["Golang"] = "Golang"

	fmt.Println(languages)
	fmt.Printf("%T\n", languages)


}
