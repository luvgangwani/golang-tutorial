package main

import (
	"fmt"
	"sort"
)

func main() {
	var testStrings = []string{"Test1", "Test2", "Test3"}

	fmt.Println("Slice values are: ", testStrings)
	fmt.Printf("Type of slice is %T\n", testStrings)

	// add new values to the slice
	testStrings = append(testStrings, "Test4", "Test5")
	fmt.Println("New slice values are: ", testStrings)

	// slice the list from 1..n
	testStrings = append(testStrings[1:])
	fmt.Println("Slice values strting the 1st index are: ", testStrings)

	// slice the list from 1..3 (3rd index not inclusive)
	testStrings = append(testStrings[1:3])
	fmt.Println("Slice values strting the 1st index to 3rd index are: ", testStrings)

	// slice the list from start to 3rd index (3rd index not inclusive)
	testStrings = []string{"Test1", "Test2", "Test3", "Test4", "Test5"}
	testStrings = append(testStrings[:3])
	fmt.Println("Slice values from the start index to 3rd index are: ", testStrings)

	// create slices using the make method

	nums := make([]int, 4)

	nums[0] = 123
	nums[1] = 345
	nums[2] = 789
	nums[3] = 567

	fmt.Println("List of numbers are", nums)

	// nums[4] = 789 // this will not work as nums are only assigned 4 memory locations

	nums = append(nums, 789, 901)

	fmt.Println("Update list of numbers are", nums)

	// Sorting slices

	fmt.Println("Is nums sorted?", sort.IntsAreSorted(nums)) // check if ints are sorted
	sort.Ints(nums)
	fmt.Println("Sorted list of nums are", nums)
	fmt.Println("Is nums sorted?", sort.IntsAreSorted(nums)) // check if ints are sorted

	// Remove values from a slice

	var programmingLanguages = []string{"C", "C++", "Python", "Rust", "Go"}

	var removeIx int = 2

	programmingLanguages = append(programmingLanguages[:removeIx], programmingLanguages[removeIx+1:]...)
	fmt.Println("Slice value after removing element at index 2 is", programmingLanguages)
}
