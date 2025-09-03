package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Reverser interface {
	Reverse() any
}

type Number struct {
	Value int `json:"value"`
}

func (n Number) Reverse() any {
	x := n.Value
	rev := 0

	for x != 0 {
		digit := x % 10
		rev = rev*10 + digit
		x = x / 10
	}

	return Number{Value: rev}
}

type ReverseHandler struct{}

func (ReverseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "use POST", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var n Number
	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}

	var rev Reverser = n
	result := rev.Reverse()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "encode error", http.StatusInternalServerError)
	}
}

func main() {
	http.Handle("/reverse", ReverseHandler{})
	fmt.Println("Сервер запущен: http://localhost:8080/reverse")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
