package main

import (
	"./auth"
	"./restapi"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()

	restapi.Books = append(restapi.Books, restapi.Book{ID: "1", Isbn: "448743", Title: "Book one", Author: &restapi.Author{Firstname: "joe", Lastname: "doe"}})
	restapi.Books = append(restapi.Books, restapi.Book{ID: "2", Isbn: "848743", Title: "Book two", Author: &restapi.Author{Firstname: "steven", Lastname: "bremen"}})

	r.HandleFunc("/api/books", auth.BasicAuthMiddleware(restapi.GetBooks)).Methods("GET")
	r.HandleFunc("/api/books/{id}", auth.BasicAuthMiddleware(restapi.GetBook)).Methods("GET")
	r.HandleFunc("/api/books", restapi.CreateBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}", restapi.UpdateBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}", restapi.DeleteBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}