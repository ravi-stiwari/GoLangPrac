package main

import (
	"fmt"
)

func main() {

	//var x float64 = -42
	var result string

	if x := -42; x < 0 {
		result = "Less than zero"
	} else if x == 0 {
		result = "Greater than to zero"
	} else {
		result = "Greater than to zero"
	}
	fmt.Println("Result :", result)
	//x is only available for if else
	//fmt.Println("Value of x :", x)
}
