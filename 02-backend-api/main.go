package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/add", http.HandlerFunc(add))
	http.Handle("/subtract", http.HandlerFunc(subtract))
	http.Handle("/multiply", http.HandlerFunc(multiply))
	http.Handle("/divide", http.HandlerFunc(divide))
	http.Handle("/sum", http.HandlerFunc(sum))

	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type InputNumbers struct {
	FirstNumber  int `json:"num1"`
	SecondNumber int `json:"num2"`
}

func add(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if req.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var i InputNumbers
	err := json.NewDecoder(req.Body).Decode(&i)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	addResult := i.FirstNumber + i.SecondNumber

	resultJSON, err := json.Marshal(addResult)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resultJSON)
}

func subtract(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if req.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var i InputNumbers
	err := json.NewDecoder(req.Body).Decode(&i)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	subtractResult := i.FirstNumber - i.SecondNumber

	resultJSON, err := json.Marshal(subtractResult)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resultJSON)
}

func multiply(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if req.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var i InputNumbers
	err := json.NewDecoder(req.Body).Decode(&i)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mulitplyResult := i.FirstNumber * i.SecondNumber

	resultJSON, err := json.Marshal(mulitplyResult)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resultJSON)
}

func divide(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if req.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var i InputNumbers
	err := json.NewDecoder(req.Body).Decode(&i)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if i.SecondNumber == 0 {
		val, err := json.Marshal("Cannot Divide by 0!!!!")
		if err != nil {
			return
		}
		w.Write(val)
		return
	}

	divideResult := i.FirstNumber / i.SecondNumber

	resultJSON, err := json.Marshal(divideResult)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resultJSON)
}

// array of variable length as input
func sum(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {

	}
	fmt.Println("sum endpoint")
}
