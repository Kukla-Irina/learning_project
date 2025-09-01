package main

import (
	"fmt"
	"log"
	"net/http"
)

type MyHandler struct{}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from custom handler!")
}

func main() {
	fmt.Println("Server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", MyHandler{}); err != nil {
		log.Fatal(err)
	}
}
