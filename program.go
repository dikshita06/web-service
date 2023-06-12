package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Numbers struct {
	FirstNumber  int `json:"first_number"`
	SecondNumber int `json:"second_number"`
}

type Result struct {
	Result int `json:"result"`
}

func sumHandler(w http.ResponseWriter, r *http.Request) {
	var numbers Numbers
	err := json.NewDecoder(r.Body).Decode(&numbers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sum := numbers.FirstNumber + numbers.SecondNumber
	result := Result{Result: sum}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Fatal(err)
	}
}

func diffHandler(w http.ResponseWriter, r *http.Request) {
	var numbers Numbers
	err := json.NewDecoder(r.Body).Decode(&numbers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	diff := numbers.FirstNumber - numbers.SecondNumber
	result := Result{Result: diff}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Fatal(err)
	}
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	var numbers Numbers
	err := json.NewDecoder(r.Body).Decode(&numbers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product := numbers.FirstNumber * numbers.SecondNumber
	result := Result{Result: product}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Fatal(err)
	}
}

func divideHandler(w http.ResponseWriter, r *http.Request) {
	var numbers Numbers
	err := json.NewDecoder(r.Body).Decode(&numbers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if numbers.SecondNumber == 0 {
		http.Error(w, "Cannot divide by zero", http.StatusBadRequest)
		return
	}

	divide := numbers.FirstNumber / numbers.SecondNumber
	result := Result{Result: divide}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Fatal(err)
	}
}

func moduloHandler(w http.ResponseWriter, r *http.Request) {
	var numbers Numbers
	err := json.NewDecoder(r.Body).Decode(&numbers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if numbers.SecondNumber == 0 {
		http.Error(w, "Cannot perform modulo with zero", http.StatusBadRequest)
		return
	}

	mod := numbers.FirstNumber % numbers.SecondNumber
	result := Result{Result: mod}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/sum", sumHandler)
	http.HandleFunc("/difference", diffHandler)
	http.HandleFunc("/product", productHandler)
	http.HandleFunc("/divide", divideHandler)
	http.HandleFunc("/modulo", moduloHandler)

	fmt.Println("Server is running on http://localhost:9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
