package backend

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func Run(addr string) {
	http.HandleFunc("/", helloWorld)
	fmt.Println("Server started and listening on port ", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
