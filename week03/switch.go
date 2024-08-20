package week03

import (
	"fmt"
)

func Switch() {
	today := 26

	switch today {
	case 5:
		fmt.Println("Today is 5th. Clean your house.")
		fallthrough
	case 10:
		fmt.Println("Today is 10th. Buy some wine.")
		fallthrough
	case 14:
		fmt.Println("Hi.")
		fallthrough
	case 15:
		fmt.Println("Today is 15th. Visit a doctor.")
	case 25, 26, 27:
		fmt.Println("Today is 25th. Buy Some food.")
		fallthrough
	case 31:
		fmt.Println("Party tonight. ")
		fallthrough
	default:
		fmt.Println("No information available for that day")
	}
}
