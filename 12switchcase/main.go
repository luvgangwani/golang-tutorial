package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	diceNum := rand.Intn(6) + 1
	fmt.Printf("Dice has rolled number %v\n", diceNum)

	switch diceNum {
	case 1:
		fmt.Println("You can move 1 spot")
	case 2:
		fmt.Println("You can move 2 spots")
	case 3:
		fmt.Println("You can move 3 spots")
		fallthrough // run this case and fall through to the next case
	case 4:
		fmt.Println("You can move 4 spots")
		fallthrough // run this case and fall through to the next case
	case 5:
		fmt.Println("You can move 5 spots")
	case 6:
		fmt.Println("You can move 6 spots and roll the dice one more time.")
	default:
		fmt.Println("Please roll the dice again.")
	}
}
