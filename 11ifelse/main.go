package main

import "fmt"

func main() {
	x := 5

	if x < 10 {
		fmt.Println("Value of x is less than 10")
	} else if x > 10 {
		fmt.Println("Value of x is more than 10")
	} else {
		fmt.Println("Value of x is equal to 10")
	}

	// declare variables in if statements

	if y := 5; y > 10 {
		fmt.Println("Value of y is greater than 10")
	} else {
		fmt.Println("Value of y is less than or equal to 10")
	}
}
