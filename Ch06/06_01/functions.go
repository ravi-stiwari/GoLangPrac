package main

import "fmt"

//Function Overloading is not there

func main() {
	doSomething()

	sum := addValues(23, 54)
	fmt.Println("Sum :", sum)

	sum = addValues1(23, 54)
	fmt.Println("Sum :", sum)

	sum = addAllValues(12, 54, 79)
	fmt.Println("Sum :", sum)
}

func doSomething() {
	fmt.Println("Doing Something")
}

func addValues(value1 int, value2 int) int {
	return value1 + value2
}

func addValues1(value1, value2 int) int {
	return value1 + value2
}

//Slice of int as args
func addAllValues(values ...int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	fmt.Printf("Type of Argument is %T\n", values)
	return sum
}
