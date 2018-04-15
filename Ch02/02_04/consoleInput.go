package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	/*
	   This will take single word
	*/
	//var name string
	//fmt.Scanln(&name)
	//fmt.Printf("Hello %v", name)

	/*
	   This will take multiple words till anyone press enter key or line feed
	*/
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text : ")
	str, _ := reader.ReadString('\n') // Single Quote is treated as a byte instead of string
	fmt.Println(str)

	/*
	   What if we need validation that user should enter a specific type of value
	*/
	fmt.Print("Enter a number : ")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Value of Number is %v", f)
	}
}
