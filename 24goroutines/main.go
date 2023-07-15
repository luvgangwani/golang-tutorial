package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex

var signals = []string{}

func main() {
	websites := []string{
		"https://google.com",
		"https://lco.dev",
		"https://go.dev",
		"https://github.com",
		"https://fb.com",
	}

	for _, website := range websites {
		go getStatusCodes(website)
		wg.Add(1)
	}
	wg.Wait() // Wait for the function to end as there are parallel processes running
	fmt.Println("Signals are", signals)
}

func getStatusCodes(website string) {
	defer wg.Done()
	res, err := http.Get(website)

	if err != nil {
		fmt.Println("Issue fetching status code for", website)
	} else {

		mut.Lock()
		signals = append(signals, website)
		mut.Unlock()
		fmt.Printf("Status code for %d: %s\n", res.StatusCode, website)
	}
}
