package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Dates and times")
	n := time.Now()
	fmt.Println("Time now is: ", n)
	fmt.Println("Time now is ANSI formating: ", n)
	t := time.Date(2024, time.February, 3, 22, 1, 3, 4, time.Local)
	fmt.Println("Date created is: ", t)
	fmt.Println("Date created is ANSI formating: ", t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Sat Feb  3 22:01:03 2024")
	fmt.Printf("Created time parsed is of type %T\n", parsedTime)

}
