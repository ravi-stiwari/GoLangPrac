package main

import (
	"fmt"
)

func main() {
	var p *int

	//panic: runtime error: invalid memory address or nil pointer dereference
	//fmt.Println("Value of p:", *p)

	if p != nil {
		fmt.Println("Value of p:", *p)
	} else {
		fmt.Println("p is nil")
	}

	var intValue = 42
	p = &intValue

	if p != nil {
		fmt.Println("Value of p:", *p)
	} else {
		fmt.Println("p is nil")
	}

	var floatValue float64
	floatValue = 42.13
	pointer := &floatValue
	fmt.Println("Float value:", *pointer)

	*pointer = *pointer / 31
	fmt.Println("Float value using pointer:", *pointer)
	fmt.Println("Float value using variable:", floatValue)
}
