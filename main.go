package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello wolrd")
}

func main() {
	http.HandleFunc("/", helloWorld)
	fmt.Println("Server started and listening on localhost:3003")
	log.Fatal(http.ListenAndServe(":3003", nil))
}
