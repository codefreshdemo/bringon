package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bring On the Builds!")
}

func BuildIndex(w http.ResponseWriter, r *http.Request) {
	builds := Builds{
		Build{Name: "#3455"},
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

func BuildAdd(w http.ResponseWriter, r *http.Request) {
	var build Build
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &build); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	log.Printf("build is %v", build)
	session := dbinit()
	//log.Printf("got collection %v", bCol)
	bCol := session.DB("bubblz").C("builds")
	log.Printf("got collection")

	t := bCol.Insert(&build)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

}
