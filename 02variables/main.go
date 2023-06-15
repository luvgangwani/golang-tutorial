package main

import "fmt"

const AuthType = "2FA" // Public variable. Can be accessed anywhere in this file and outside.

func main() {
	var username string = "test-username"
	fmt.Println(username)
	fmt.Printf("Variable is of type %T \n", username)

	var isSelected bool = true
	fmt.Println(isSelected)
	fmt.Printf("Variable is of type %T \n", isSelected)

	var smallInt uint8 = 255
	fmt.Println(smallInt)
	fmt.Printf("Variable is of type %T \n", smallInt)

	var smallFloat float32 = 255.235754543456865
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type %T \n", smallFloat)

	var largeFloat float64 = 255.235754543456865
	fmt.Println(largeFloat)
	fmt.Printf("Variable is of type %T \n", largeFloat)

	// default value and aliases
	var num int
	fmt.Println(num)
	fmt.Printf("Variable is of type %T \n", num)

	// implicit type
	var foo = "bar"
	fmt.Println(foo)
	fmt.Printf("Variable is of foo %T \n", foo)

	// no var keyword
	// := is a walrus operator. This cannot be used outside a function
	numberOfIterations := 5
	fmt.Println(numberOfIterations)
	fmt.Printf("Variable is of type %T \n", numberOfIterations)

	fmt.Println(AuthType)
	fmt.Printf("Variable is of type %T \n", AuthType)
}
