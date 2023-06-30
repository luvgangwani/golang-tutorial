package main

import "fmt"

func main() {

	var ptr *int
	fmt.Println("Value of the pointer is", ptr) // return <nil> as no value is assigned

	myNum := 33

	var ptrRef = &myNum // reference to myNum
	fmt.Println("Reference of myNum is", ptrRef)
	fmt.Println("Value at the reference is", *ptrRef)

	// Performing operations on a value stored at a reference
	*ptrRef = *ptrRef + 2
	fmt.Println("Value at the reference is", *ptrRef)
	fmt.Printf("Value of myNum is also %d as it points to the same memory address\n", myNum)
}
