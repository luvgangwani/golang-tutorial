package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Mod in golang")
	route := mux.NewRouter()
	route.HandleFunc("/", serveHome)

	log.Fatal(http.ListenAndServe(":4000", route))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Response from Golang server</h1>"))
}
