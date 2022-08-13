package main

import (
	"github.com/Yash-23/crud/pkg/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books",handlers.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}",handlers.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books",handlers.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}",handlers.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}",handlers.DeleteBook).Methods(http.MethodDelete)

	log.Println("API is running")
	http.ListenAndServe(":8080", router)
}
