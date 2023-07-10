package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, err := http.Get("http://lco.dev")

	if err != nil {
		panic(err)
	}

	fmt.Printf("Type of response is %T\n", response)

	defer response.Body.Close()

	byteResponse, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(byteResponse))
}
