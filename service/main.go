package main

import (
	"log"
	"net/http"
	"strconv"

	bringon "github.com/antweiss/bringon"
)

func main() {

	router := bringon.NewRouter()
	port := 8091
	log.Println("Listening on port ", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
