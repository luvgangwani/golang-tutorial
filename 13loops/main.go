package main

import "fmt"

func main() {
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	// general for loop
	fmt.Println("List of all months using a general for loop")
	for month := 0; month < len(months); month++ {
		fmt.Println(months[month])
	}

	// for using range
	fmt.Println("List of all months using for..range loop")
	for month := range months {
		fmt.Println(months[month])
	}

	// for each variant
	fmt.Println("List of all months using a for each variant")
	for ix, month := range months {
		fmt.Printf("Month %v: %v\n", ix+1, month)
	}

	// while variant
	fmt.Println("List of all months using a while variant")
	i := 0

	for i < len(months) {
		fmt.Println(months[i])
		i++
	}

	// break statement
	fmt.Println("List of all months using a break statement")
	j := 0

	for j < len(months) {
		if j == 6 {
			break
		}
		fmt.Println(months[j])
		j++
	}

	// continue statement
	fmt.Println("List of all months using a continue statement")
	k := 0

	for k < len(months) {
		if k == 6 {
			k++
			continue
		}
		fmt.Println(months[k])
		k++
	}

	// goto statement
	fmt.Println("List of all months using a goto statement")
	l := 0

	for l < len(months) {
		if l == 6 {
			goto july
		}
		fmt.Println(months[l])
		l++
	}

july:
	fmt.Println("Came here from a goto statement")
}
