package main

import (
	"encoding/json"
	"net/http"
)

type Person struct {
	Name string
	Age  int8
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	alexander := Person{"Alexander", 22}
	js, err := json.Marshal(alexander)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
