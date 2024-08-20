package week02

import (
	"fmt"
)

func Variables() {
	// Hello world
	fmt.Println(".......week02......")
	fmt.Println("Hello World")
	// Variables
	// Explicit type declaration
	var var1 string = "Hello world"
	//Implicit
	var2 := "Hi"

	fmt.Println(var1, var2)

	//Just declaration
	var (
		var3 string
		var4 int
		var5 bool
	)
	fmt.Println(var3, var4, var5) //<- no initialized variables

	// Formated print
	fmt.Printf("%T\n", var3) //<- print type
	fmt.Printf("%v\n", var3) // <- print value

	var boolean bool
	var float float32
	var cad string

	fmt.Println(boolean, float, cad)

	// Constants
	const (
		i int    = 1
		j int    = 20
		k string = "adiós"
	)
	fmt.Println(i, j, k)
}
