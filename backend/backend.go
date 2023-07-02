package backend

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func Run(addr string) {
	http.HandleFunc("/", helloWorld)
	fmt.Println("Server running on port :8090")
	http.ListenAndServe(addr, nil)
}
