package main

import "fmt"

func main() {
	units := make(map[string]string)

	units["km"] = "kilometer"
	units["m"] = "meter"
	units["cm"] = "centimeter"

	fmt.Println("Units details are", units)

	// add a new value
	units["l"] = "litre"
	fmt.Println("Updated units details are", units)

	// delete a value
	delete(units, "cm")
	fmt.Println("Updated units details after deletion are", units)

	fmt.Println("===========")

	// loop through a map
	for key, value := range units {
		fmt.Printf("For key %v, the value is %v\n", key, value)
	}

	fmt.Println("===========")

	// loop through a map using the comma/ok syntax
	fmt.Println("The values of the map are:")
	for _, value := range units {
		fmt.Printf("%v\n", value)
	}
}
