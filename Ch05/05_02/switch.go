package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	// dow := rand.Intn(6) + 1
	//fmt.Println("Day", dow)

	result := ""

	switch dow := rand.Intn(6) + 1; dow {
	case 1:
		result = "Sunday"
	case 7:
		result = "Saturday"
	default:
		result = "Weekday"
	}

	//fmt.Println("Day", dow, " :", result)
	fmt.Println("Day", ":", result)

	x := -42

	//Break keyword is not mandatory

	switch {
	case x < 0:
		result = "Less than zero"
		//This is used to get into just next case
		//fallthrough
	case x > 0:
		result = "Greater than zero"
	case x == 0:
		result = "Equal to zero"
	}
	fmt.Println("Value of x :", x, "Result :", result)
}
