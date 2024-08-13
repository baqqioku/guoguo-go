package day01

import (
	"fmt"
)

func Max(num1, num2 int) int {
	fmt.Println(num1, num2)
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}
