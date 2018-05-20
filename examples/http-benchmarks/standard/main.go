package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	log.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", http.HandlerFunc(handler)))
}

func handler(w http.ResponseWriter, r *http.Request) {
	var body Body
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(body)
	if err != nil {
		panic(err)
	}
}

type Body struct {
	Colors []Color `json:"colors"`
}

type Color struct {
	Color    string `json:"color,omitempty"`
	Category string `json:"category,omitempty"`
	Type     string `json:"type,omitempty"`
	Code     Code   `json:"code,omitempty"`
}

type Code struct {
	RGBA []int  `json:"rgba,omitempty"`
	Hex  string `json:"hex,omitempty"`
}
