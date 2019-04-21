package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/protected", protected).Methods("GET")

	var port = ":6000"
	fmt.Println("\n*** API running on http://localhost" + port + " ***\n")
	log.Fatal(http.ListenAndServe(port, handlers.LoggingHandler(os.Stdout, r)))
}

var home = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Sanity</h1>")
}

var protected = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
