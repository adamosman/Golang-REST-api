package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

// Get all Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// Loop through books and find with ID
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create a New Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book // variable 'book' set to the Book struct
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID - NOT SAFE
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book // variable 'book' set to the Book struct
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

// Delete Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

// Upload for files
func Upload(w http.ResponseWriter, r *http.Request) {
	var str string
	counts := map[string]int{}
	if r.Method == http.MethodPost {
		f, _, err := r.FormFile("usrfile")
		if err != nil {
			log.Println(err)
			http.Error(w, "Error uploading file", http.StatusInternalServerError)
			return
		}
		defer f.Close()

		bs, err := ioutil.ReadAll(f)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}
		str = string(bs)
	}
	for _, word := range strings.Fields(str) {
		counts[word]++
	}
	json, _ := json.MarshalIndent(counts, "", " ")
	words := strings.Fields(str)
	w.Header().Set("CONTENT-TYPE", "text/html; charset=UTF-8")
	fmt.Fprintf(w, `<form action="/" method="post" enctype="multipart/form-data">
					Upload a file<br>
					<input type="file" name="usrfile"><br>
					<input type="submit">
					</form>
					<br>
					<br>
					<h1>Unique word count: %s</h1>
					<br>
					<h1>Total word count: %v</h1>`, json, len(words))
}

// WordCount counts the frequency of each word
func WordCount(str string) map[string]int {
	counts := map[string]int{}
	for _, word := range strings.Fields(str) {
		counts[word]++
	}
	return counts
}

// TotalWordCount counts total number of words
func TotalWordCount(s string) int {
	words := strings.Fields(s)
	return len(words)
}

func main() {
	// Init Router
	// r := mux.NewRouter()

	// Mock Data - @todo implement DB
	// books = append(books, Book{ID: "1", Isbn: "526421", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	// books = append(books, Book{ID: "2", Isbn: "364213", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	// Route Handlers / Endpoints
	// r.HandleFunc("/api/books", getBooks).Methods("GET")
	// r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	// r.HandleFunc("/api/books", createBook).Methods("POST")
	// r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	// r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// log.Fatal(http.ListenAndServe(":8000", r))
	http.HandleFunc("/", Upload)
	http.ListenAndServe(":8000", nil)
}
