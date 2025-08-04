package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("15:04:05")
	fmt.Fprintf(w, "Текущее время: %s", currentTime)
}

func main() {
	http.HandleFunc("/time", handler)
	fmt.Println("Сервер запущен через порт 8080.")
	http.ListenAndServe(":8080", nil)
}
