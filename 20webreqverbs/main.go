package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("GET Request")

	GetReq()
	PostReq()
}

func GetReq() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Response status", response.StatusCode)
	fmt.Println("Response content length", response.ContentLength)

	var stringBuilder strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := stringBuilder.Write(content)

	fmt.Println("Byte count is", byteCount)

	fmt.Println(stringBuilder.String())

	// fmt.Println(string(content))
}

func PostReq() {

	data := strings.NewReader(`
		{
			"foo1": "bar1",
			"foo2": "bar2"
		}
	`)
	response, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
