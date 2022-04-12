package main

import (
	"github.com/gorilla/mux"
	"lesson_2/utils"
	"log"
	"net/http"
)

const (
	API_PREFIX            = "/api/v1"
	PORT                  = "8000"
	BOOK_RESOURCE_PREFIX  = API_PREFIX + "/book"
	BOOKS_RESOURCE_PREFIX = API_PREFIX + "/books"
)

func main() {
	log.Println("Server is starting")

	router := mux.NewRouter()
	utils.BuildBookResource(router, BOOK_RESOURCE_PREFIX)
	utils.BuildBooksResource(router, BOOKS_RESOURCE_PREFIX)
	log.Fatalln(http.ListenAndServe(":"+PORT, router))
}
