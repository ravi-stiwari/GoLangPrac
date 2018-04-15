package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go Programming Language Launched on: %v\n", t)

	now := time.Now()
	fmt.Printf("Time now is : %v\n", now)

	fmt.Println("Month is", t.Month())
	fmt.Println("Day is", t.Day())
	fmt.Println("Weekday is", t.Weekday())

	tomorrow := t.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v, %v, %v, %v.\n", tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())

	// Custom Date time Format
	longFormat := "Monday, January, 2, 2006"
	fmt.Println("Tomorrow is ", tomorrow.Format(longFormat))

	shortFormat := "1/2/06"
	fmt.Println("Tomorrow is ", tomorrow.Format(shortFormat))

	//Using available date format constant
	fmt.Println("Tomorrow is ", tomorrow.Format(time.UnixDate))
}
