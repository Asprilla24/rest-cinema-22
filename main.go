package main

import (
	"log"
	"net/http"

	"github.com/Asprilla24/rest-cinema-22/store"
)

func main() {
	router := store.NewRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}
