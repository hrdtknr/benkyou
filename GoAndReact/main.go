package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Fruit struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

var fruits = []Fruit{
	{ID: 1, Name: "apple", Icon: "🍎"},
	{ID: 2, Name: "banana", Icon: "🍌"},
	{ID: 3, Name: "grape", Icon: "🍇"},
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello Worldddddddddddddd")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(fruits)
}

func main() {
	fmt.Println("sss")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
