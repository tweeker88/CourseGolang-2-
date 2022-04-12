package utils

import (
	"github.com/gorilla/mux"
	"lesson_2/handlers"
)

func BuildBookResource(router *mux.Router, prefix string) {
	router.HandleFunc(prefix+"/{id}", handlers.GetBookById).Methods("GET")       // Получение книги
	router.HandleFunc(prefix, handlers.CreateBook).Methods("POST")               // Добавление книги
	router.HandleFunc(prefix+"/{id}", handlers.UpdateBookById).Methods("PUT")    // Изменение книги
	router.HandleFunc(prefix+"/{id}", handlers.DeleteBookById).Methods("DELETE") // Удаление книги
}

func BuildBooksResource(router *mux.Router, prefix string) {
	router.HandleFunc(prefix, handlers.GetBooks).Methods("GET") // Получение всех книг
}
