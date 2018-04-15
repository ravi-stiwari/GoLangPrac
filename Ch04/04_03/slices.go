package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	colors = append(colors, "Purple")
	fmt.Println(colors)

	//Delete entry from a slice
	//len(colors) is default
	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	colors = append(colors[1:])
	fmt.Println(colors)

	//Delete last item
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	//Make(Init Obj, Size, Capacity)
	numbers := make([]int, 5, 5)
	numbers[0] = 134
	numbers[1] = 72
	numbers[2] = 32
	numbers[3] = 12
	numbers[4] = 156

	fmt.Println(numbers)

	numbers = append(numbers, 235)
	fmt.Println(numbers)
	//If anything is beyond capacity it will add number of new places for future use
	fmt.Println(cap(numbers))

	//Sorting Slices
	sort.Ints(numbers)
	fmt.Println(numbers)
}
