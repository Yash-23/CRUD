package handlers

import (
	"encoding/json"
	"github.com/Yash-23/crud/pkg/mocks"
	"github.com/Yash-23/crud/pkg/models"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}
	var updatedBook models.Book
	json.Unmarshal(body, &updatedBook)

	for index, book := range mocks.Books{
		if book.Id == id{
			book.Title = updatedBook.Title
			book.Author = updatedBook.Author
			book.Desc = updatedBook.Desc

			mocks.Books[index ] = book

			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode("Updated" )
			break
		}
	}

}

