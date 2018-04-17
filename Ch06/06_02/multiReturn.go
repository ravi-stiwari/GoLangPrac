package main

import (
	"fmt"
)

func main() {
	name, length := FullName("Ravi Shanker", "Tiwari")
	fmt.Printf("Fullname: %v, number of chars: %v\n", name, length)

	name1, length1 := FullName("Ankit", "Shukla")
	fmt.Printf("Fullname: %v, number of chars: %v\n", name1, length1)
}

//FullName function is for getting the full name and char count by passing first and lastname
func FullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)
	return full, length
}

//FullNameNakedReturn function is for getting the full name and char count by passing first and lastname
func FullNameNakedReturn(f, l string) (full string, length int) {
	full = f + " " + l
	length = len(full)
	return
}
