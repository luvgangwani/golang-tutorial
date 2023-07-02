package main

import "fmt"

func main() {
	newUser := User{FirstName: "Test", LastName: "Name", Status: true, Username: "testname", Age: 24}

	fmt.Println("New user details are", newUser)

	fmt.Printf("All details in key-value pairs are %+v\n", newUser)

	fmt.Printf("Name: %v; Username: %v\n", newUser.FirstName+" "+newUser.LastName, newUser.Username)
}

type User struct {
	FirstName string
	LastName  string
	Status    bool
	Username  string
	Age       int
}
