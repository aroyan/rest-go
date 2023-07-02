package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func main() {
	http.HandleFunc("/", helloWorld)
	fmt.Println("Server is running on port 8090")
	http.ListenAndServe(":8090", nil)
}
