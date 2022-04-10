package main

import (
	"fmt"
	"log"
	"net/http"
)

func GetGreet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hi! I'm new web-server!</h1>")

}

func RequestHandler() {
	http.HandleFunc("/", GetGreet)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	RequestHandler()
}
