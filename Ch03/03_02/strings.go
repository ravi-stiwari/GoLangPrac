package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "An Implicit Typed String"
	fmt.Printf("str value is %v and of datatype %T\n", str, str)

	var str2 string
	str2 = "An Explicit Typed String"
	fmt.Printf("str value is %v and of datatype %T\n", str2, str2)

	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.Title(str))

	lvalue := "hello"
	uvalue := "HELLO"
	fmt.Println("Equals ? ", lvalue == uvalue)

	//Non Case sensitive comparision
	fmt.Println("Equals Non Case Sensitive? ", strings.EqualFold(lvalue, uvalue))

	//Find String inside string and it is case sensitive as well
	fmt.Println("Contains exp ? ", strings.Contains(str2, "Exp"))
	fmt.Println("Contains exp ? ", strings.Contains(str, "Exp"))
	fmt.Println("Contains exp ? ", strings.Contains(str2, "exp"))
}
