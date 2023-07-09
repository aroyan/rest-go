package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getReq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is GET method")
}

func postReq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is POST method")
}

func patchReq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is PATCH method")
}

func deleteReq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is DELETE method")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", getReq).Methods("GET")
	r.HandleFunc("/", postReq).Methods("POST")
	r.HandleFunc("/", patchReq).Methods("PATCH")
	r.HandleFunc("/", deleteReq).Methods("DELETE")

	http.Handle("/", r)
	fmt.Println("Server is running on port :8009")
	log.Fatal(http.ListenAndServe(":8009", nil))
}
