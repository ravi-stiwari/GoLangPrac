package main

import (
	"fmt"
)

//Dog is to be used when an Dog object need to be used
type Dog struct {
	Breed  string
	Weight int
}

func main() {
	poodle := Dog{"Poodle", 37}
	fmt.Println(poodle)

	// %+v is used to get described value of an object
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
}
