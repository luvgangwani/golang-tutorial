package main

import "fmt"

func main() {
	var carList [4]string // the length should always be mentioned

	carList[0] = "BMW"
	carList[1] = "Audi"
	carList[3] = "Nissan"

	fmt.Println("Array with car list:", carList)                      // second index will be represented by a space
	fmt.Println("Length of the array with car list is", len(carList)) // The length will always be the same as the one specified during initialization

	var bikeList = [3]string{"Kawasaki", "Honda", "Yamaha"}

	fmt.Println("Array with bike list:", bikeList)
	fmt.Println("Length of the array with bike list is", len(bikeList)) // The length will always be the same as the one specified during initialization
}
