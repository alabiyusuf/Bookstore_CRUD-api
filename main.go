package main

import (
	"encoding/json"
	"log"
	"net/http"

	"math/rand"
	"strconv"

	"github.com/gorilla/mux"
)

// CREATING BOOKS STRUCTS (MODEL)

type Book struct {
	ID		string	`json:"id"`
	Isbn	string	`json:"isbn"`
	Title	string	`json:"title"`
	Author	*Author	`json:"author"`
}

//  AUTHOR STRUCT

type Author struct {
	FirstName	string	`json:"firstname"`
	LastName	string	`json:"lastname"`
}


//  BOOKS STRUCT AS A SLICE 	
var books []Book

// GET ALL BOOKS
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json ")	
	json.NewEncoder(w).Encode(books)
}

// GET A SINGLE BOOK 
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json ")

	// GET THE PARAMS
	params := mux.Vars(r)

	// LOOP THROUGH THE BOOKS TO FIND THE ID
	for _, item := range books {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// CREATE A NEW BOOK
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json ")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	// CREATING A MOCK ID
	book.ID = strconv.Itoa(rand.Intn(100000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)


}

// UPDATE BOOK
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item	:= range books {
		if item.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)
			var book Book
		_ = json.NewDecoder(r.Body).Decode(&book)

		book.ID = params["id"]
		books = append(books, book)
		json.NewEncoder(w).Encode(book)
		return
		}
	}
	json.NewEncoder(w).Encode(books)
}

// DELETE BOOK
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item	:= range books {
		if item.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}



func main () {
	// INIT ROUTER
	r := mux.NewRouter()

	// MOCK DATA FOR IMPLEMENTING DATABASE

	books = append(books, Book{ID: "1", Isbn: "448734", Title: "Understanding Terminologies V1", Author: &Author{FirstName: "John", LastName: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "448735", Title: "Understanding Terminologies V2", Author: &Author{FirstName: "Smilga", LastName: "Shaun"}})

	// ROUTER HANDLERS
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE") 

	// RUN SERVER
	log.Fatal(http.ListenAndServe(":8000", r))

}