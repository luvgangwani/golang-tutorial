package main

import (
	"fmt"
	"net/url"
)

const testUrl string = "https://lco.dev:5122/learn?foo1=bar1&foo2=bar2"

func main() {
	result, _ := url.Parse(testUrl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	queryParams := result.Query()

	fmt.Printf("Type of query params is %T\n", queryParams)

	fmt.Println(queryParams["foo1"])
	fmt.Println(queryParams["foo2"])

	urlParts := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "foo=bar",
	}

	constructedURL := urlParts.String()
	fmt.Println(constructedURL)

}
