package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/meghsha/g--bookstore/pkg/models"
	"github.com/meghsha/g--bookstore/pkg/utils"
)

func getBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.statusOK)
	w.Write(res)
}

func getBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookid := vars["bookId"]
	ID, err := strconv.ParseInt(bookid, 0, 0)
	if err != nil {
		fmt.Println("Failed to parse the ID")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.statusOK)
	w.Write(res)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	book := CreateBook.CreateBook()
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.statusOK)
	w.Write(res)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Failed to parse the ID")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.statusOK)
	w.Write(res)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var updateBookDetails = &models.Book{}
	utils.ParseBody(r, updateBookDetails)

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Failed to parse the ID")
	}
	bookDetails, db := models.GetBookById(ID)

	if updateBookDetails.Name != "" {
		bookDetails.Name = updateBookDetails.Name
	}
	if updateBookDetails.Author != "" {
		bookDetails.Author = updateBookDetails.Author
	}
	if updateBookDetails.Publisher != "" {
		bookDetails.Publisher = updateBookDetails.Publisher
	}

	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.statusOK)
	w.Write(res)
}
