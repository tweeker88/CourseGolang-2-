package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"lesson_2/models"
	"log"
	"net/http"
	"strconv"
)

func GetBookById(writer http.ResponseWriter, r *http.Request) {
	initHeaders(writer)

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("Error! Can't convert query parameter", err)
	}

	log.Println("Get info about book with id", id)

	book, found := models.GetBookById(id)
	if found != true {
		log.Println("Book not found")
		message := models.ResponseMessage{Message: "Book not found"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(message)

		return
	}

	writer.WriteHeader(400)
	json.NewEncoder(writer).Encode(book)
}

func CreateBook(writer http.ResponseWriter, r *http.Request) {
	initHeaders(writer)

	log.Println("Create new book")

	var newBook models.Book
	errDecodeRequest := json.NewDecoder(r.Body).Decode(&newBook)
	if errDecodeRequest != nil {
		log.Println("Decode error", errDecodeRequest)
		msg := models.ResponseMessage{Message: "provideed json file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	newBook.Id = len(models.Db) + 1
	models.Db = append(models.Db, newBook)

	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(newBook)
}

func UpdateBookById(writer http.ResponseWriter, r *http.Request) {
	initHeaders(writer)

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("Error! Can't convert query parameter", err)
	}

	log.Println("Update book with id", id)

	oldBook, found := models.GetBookById(id)
	if found != true {
		log.Println("Book not found")
		message := models.ResponseMessage{Message: "Book not found"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(message)

		return
	}

	var newBook models.Book

	errDecodeRequest := json.NewDecoder(r.Body).Decode(&newBook)
	if errDecodeRequest != nil {
		log.Println("Decode error", err)
		msg := models.ResponseMessage{Message: "provideed json file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	oldBook.Title = newBook.Title
	oldBook.CreatedYear = newBook.CreatedYear
	oldBook.Author = newBook.Author
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(oldBook)
}

func DeleteBookById(writer http.ResponseWriter, r *http.Request) {
	initHeaders(writer)

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("Error! Can't convert query parameter", err)
	}
	log.Println("Delete book with Id", id)

	book, found := models.GetBookById(id)
	if found != true {
		log.Println("Book not found")
		message := models.ResponseMessage{Message: "Book not found"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(message)

		return
	}

	for key, bookInDb := range models.Db {
		if bookInDb.Id == book.Id {
			models.Db = append(models.Db[:key], models.Db[key+1:]...)
		}
	}

	msg := models.ResponseMessage{Message: "successfully deleted requested item"}
	json.NewEncoder(writer).Encode(msg)
}

func GetBooks(writer http.ResponseWriter, r *http.Request) {
	initHeaders(writer)
	log.Println("Get info about all books in database")

	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(models.Db)
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}
