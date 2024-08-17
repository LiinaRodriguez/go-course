package week02

import (
	"fmt"
	"math"
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

func AreasPerimeters() {
	fmt.Println("Areas_Perimeters...")
	var area, perimeter float32

	fmt.Println("\n Square ...............")
	var side1 float32 = 4
	perimeter = side1 * 4
	area = side1 * side1
	fmt.Printf("Side: %f cm\n", side1)
	fmt.Printf("Perimeter: %f cm\n", perimeter)
	fmt.Printf("Area: %f cm\n", area)

	fmt.Println("\nRectangle ............")
	side1 = 2
	var side2 float32 = 3

	perimeter = side1*2 + side2*2
	area = side1 * side2
	fmt.Printf("Sides, side1: %f Side2: %f \n", side1, side2)
	fmt.Printf("Perimeter: %f cm\n", perimeter)
	fmt.Printf("Area: %f cm\n", area)

	fmt.Println("\nTriangle .............")
	var (
		height float32 = 4
		base   float32 = 3
	)
	side1, side2 = 3, 3

	perimeter = base + side1 + side2
	area = (height * base) / 2
	fmt.Printf("Height: %f cm Base: %f cm\n", height, base)
	fmt.Printf("Perimeter: %f cm\n", perimeter)
	fmt.Printf("Area: %f cm\n", area)

	fmt.Println("\nRhombus .................")
	var (
		d1 float32 = 4
		D1 float32 = 8
	)
	side1 = 5
	perimeter = side1 * 4
	area = (D1 * d1) / 2
	fmt.Printf("Perimeter: %f cm\n", perimeter)
	fmt.Printf("Area: %f cm\n", area)

	fmt.Println("\nPentagon ............")
	var apothem float32 = 4
	side1 = 4
	perimeter = side1 * 5
	area = (perimeter * apothem) / 2
	fmt.Printf("Perimeter: %f cm\n", perimeter)
	fmt.Printf("Area: %f cm\n", area)

	fmt.Println("\nCircle ..................")
	var radius float32 = 3
	perimeter = 2 * radius * math.Pi
	area = math.Pi * radius * radius
	fmt.Printf("Perimeter: %f cm\n", perimeter)
	fmt.Printf("Area: %f cm\n", area)

}
