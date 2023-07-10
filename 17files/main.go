package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Create a new file")

	filename := "./test-file.txt"
	file, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println("Insert content to the new file")
	length, err := io.WriteString(file, "This is the test string I'm using for testing file handling")

	if err != nil {
		panic(err)
	}

	fmt.Println("Length of the string written to the file is", length)

	defer file.Close()

	databyte, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println("String read from the file is:", string(databyte))
}
