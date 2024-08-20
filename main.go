package main

import (
	"fmt"
	"go-course/week02"
	"go-course/week03"
)

func main() {

	s := "gopher"
	fmt.Printf("Hello and welcome, %s! \n\n", s)

	fmt.Println("Trabajo Figuras Go")
	week02.AreasPerimeters()

	fmt.Println("\n\n")
	week03.Switch()
	week03.Palindrome("rayar")
}
