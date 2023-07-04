package main

import "fmt"

func main() {
	fmt.Println("This is the main function.")
	hello()
	fmt.Println("Sum is", sum(1, 2))
	fmt.Println("Sum of n numbers is", sumN(2, 4, 6, 8, 10))

	message, result := multiply(4, 5)

	fmt.Println(message, result)
}

func hello() {
	fmt.Println("Hello")
}

func sum(numOne int, numTwo int) int {
	return numOne + numTwo
}

func sumN(nums ...int) int {
	result := 0

	for _, num := range nums {
		result += num
	}

	return result
}

func multiply(numOne int, numTwo int) (string, int) {
	return "Product of the numbers is", numOne * numTwo
}
