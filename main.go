package main

import (
	"net/http"
	"log"
)

func main() {
	router := NewRouters()
	log.Fatal(http.ListenAndServe(":8080", router))
}


