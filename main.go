package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {

	router := NewRouter()
	port := 8091
	log.Println("Listening on port ", port )
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
