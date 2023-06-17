package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your feedback on your recent purchase. (1 - 5):")

	input, _ := reader.ReadString('\n')

	fmt.Println("Thank you for submitting your rating as a", input)

	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("New result upon adding 1 to the rating is", num+1)
	}
}
