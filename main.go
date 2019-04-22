package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/protected", protected).Methods("GET")

	var port = ":6000"
	fmt.Println("\n*** API running on http://localhost" + port + " ***\n")
	log.Fatal(http.ListenAndServe(port, handlers.LoggingHandler(os.Stdout, r)))
}

var protected = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

type User struct {
	ID                              int
	Name, Email, Username, Password string
}

type Post struct {
	ID       int    `json:"id"`
	Username string `json:"Username"`
	Caption  string `json:"Caption"`
}

var posts = []Post{
	Post{ID: 1, Username: "test123", Caption: "Apples"},
	Post{ID: 2, Username: "test123", Caption: "Oranges"},
}
