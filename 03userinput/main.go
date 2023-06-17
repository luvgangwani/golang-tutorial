package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to our feedback system.")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("How was your experience with the recent product purchase? (0-5): ")

	// comma ok | err err syntax

	// read the input till a new line character is input
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for submitting your feedback with a rating", input)
	fmt.Printf("Type of input is %T \n", input)
}
