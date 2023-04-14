package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nwochaadim/go-bookstore/pkg/models"
	"github.com/nwochaadim/go-bookstore/pkg/utils"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}

	utils.ParseBody(r, book)

	b := book.CreateBook()
	res, _ := json.Marshal(b)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()

	res, _ := json.Marshal(books) // convert books to json

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.ParseInt(params["id"], 0, 0)
	book, _ := models.GetBookById(id)
	w.Header().Set("Content-Type", "application/json")

	res, _ := json.Marshal(book)

	if book.ID != uint(id) {
		w.WriteHeader(http.StatusNotFound)
		w.Write(nil)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)

	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	book.ID = uint(id)
	b := book.UpdateBook()

	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(id)

	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

	w.Write(res)
}
