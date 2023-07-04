package main

import "fmt"

type User struct {
	Email  string
	Status bool
}

func (u User) GetStatus() bool {
	return u.Status
}

func (u User) SetEmail() {
	u.Email = "updated.email@test.com"
	fmt.Println("Updated email is", u.Email)
}

func main() {
	user := User{"email@test.com", true}

	fmt.Println(user)

	fmt.Println("Status:", user.GetStatus())

	user.SetEmail()

	fmt.Println(user) // the email is not updated here as a copy of the object is passed to the method.
}
