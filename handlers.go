package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func BuildIndex(w http.ResponseWriter, r *http.Request) {
	builds := Builds{
		Build{Name: "#345"},
		Build{Name: "#3456"},
	}

	if err := json.NewEncoder(w).Encode(builds); err != nil {
		panic(err)
	}
}

func BuildShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	buildId := vars["buildId"]
	fmt.Fprintln(w, "Build show:", buildId)
}
