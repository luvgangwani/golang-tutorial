package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`                  // remove this from the resultant json
	Tags     []string `json:"keywords,omitempty"` // remove this from the resultant json if empty
}

func main() {
	courses := []Course{
		{"React JS", 2000, "reactjs.org", "reactjs", []string{"frontend", "javascript", "library"}},
		{"Python", 4000, "python.org", "python", []string{"programming", "backend"}},
		{"PHP", 1500, "php.net", "php", nil},
	}

	// resultJson, err := json.Marshal(courses)
	resultJson, err := json.MarshalIndent(courses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", resultJson)
}
