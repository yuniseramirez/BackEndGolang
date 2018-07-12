package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/gorilla/handlers"

)

type  Article struct {
	Title     string   `json:"title, omitempty"`
	Image     string   `json: "image, omitempty"`
 	Author    string   `json:"author, omitempty"`
	Paragrath string   `json:"paragrath, omitempty"`
}



var articles []Article  //yunise- maybe instead of appending the articles in the main function, you can add them to this array 



func GetArti(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(articles)
}


func main() {
	router := mux.NewRouter()
	articles = append(articles, Article{ Title: "This is the title #1", Image: "I will figure #1", Author:"Vanessa Valoy #1", Paragrath:"This is a paragrath #1"})
	articles = append(articles, Article{ Title: "This is the title #2", Image: "I will figure #2", Author:"Vanessa Valoy #2", Paragrath:"This is a paragrath #2"})
 	articles = append(articles, Article{ Title: "This is the title #3", Image: "I will figure #3", Author:"Vanessa Valoy #3", Paragrath:"This is a paragrath #3"})


	router.HandleFunc("/articles", GetArti).Methods("GET")
  	log.Print("localhost:8000")
	corsObj := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(corsObj)(router)))

}
