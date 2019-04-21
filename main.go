package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Sanity</h1>")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", home).Methods("GET")

	var port = ":6000"
	fmt.Println("\n*** API running on http://localhost" + port + " ***\n")
	log.Fatal(http.ListenAndServe(port, r))
}
