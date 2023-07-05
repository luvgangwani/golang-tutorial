package main

import "fmt"

func main() {
	/* defer fmt.Println("Hello")
	fmt.Println("World") */
	/*
		Output:
		World
		Hello
	*/

	/* defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	fmt.Println("Hello") */
	/*
		Output:
		Hello
		Three
		Two
		One
		World
	*/

	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	fmt.Println("Hello")
	testDeferFunc()

	/*
		Output:
		Hello
		4
		3
		2
		1
		0
		Three
		Two
		One
		World
	*/
}

func testDeferFunc() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
