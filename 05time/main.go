package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()

	fmt.Println("Current time:")
	fmt.Println(currentTime)

	fmt.Println("Formatted current time:")
	fmt.Println(currentTime.Format("02/01/2006 15:04:05 Mon"))

	createdDate := time.Date(2024, time.March, 5, 5, 55, 59, 0, time.UTC)
	fmt.Println("Created date:")
	fmt.Println(createdDate)

	fmt.Println("Formatted created date:")
	fmt.Println(createdDate.Format("01-02-2006 15.04.05 Monday"))
}
