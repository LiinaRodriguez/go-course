package week03

import "fmt"

func Palindrome(str string) {
	for i, le := range str {
		fmt.Printf("%v \t %c \n", i, le)

	}
}
