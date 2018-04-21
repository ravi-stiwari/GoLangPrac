package main

import (
	"fmt"
)

// func main() {
// 	//Defer is Last in First Out
// 	defer fmt.Println("Close the file")
// 	fmt.Println("Open the file")

// 	defer fmt.Println("Statement 1")
// 	defer fmt.Println("Statement 2")

// 	myFunc()

// 	defer fmt.Println("Statement 3")
// 	defer fmt.Println("Statement 4")
// 	fmt.Println("Undefered Statement")

// 	x := 1000
// 	defer fmt.Println("Value of x : ", x)
// 	x++
// 	fmt.Println("Value of x : non defer ", x)
// }

func myFunc() {
	defer fmt.Println("Defered in Func")
	fmt.Println("Not Defered in Func")
}
