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

	// Encoding
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

	// Decoding
	sourceJson := []byte(`
	{
                "coursename": "React JS",
                "Price": 2000,
                "website": "reactjs.org",
                "keywords": [
                        "frontend",
                        "javascript",
                        "library"
                ]
        }
	`)

	var source Course

	checkValid := json.Valid(sourceJson)

	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(sourceJson, &source) // pass a reference to the output variable and not a copy
		fmt.Printf("%#v\n", source)
	} else {
		fmt.Println("JSON is invalid")
	}

	// source json to a key-value pair
	// this makes use of the struct aliases (converts the JSON as is rather than moulding it into a struct)
	var mapSourceJson map[string]interface{}
	json.Unmarshal(sourceJson, &mapSourceJson)
	fmt.Printf("%#v\n", mapSourceJson)

	// Print key-value

	for k, v := range mapSourceJson {
		fmt.Printf("Key is %v, Value is %v, Type of value is %T\n", k, v, v)
	}
}
