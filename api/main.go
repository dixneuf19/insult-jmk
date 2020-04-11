package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/dixneuf19/insult-jmk/insulter"
)

var insults = insulter.CreateInsultDict()

func main() {

	router := mux.NewRouter()

	router.HandleFunc(
		"/",
		Index,
	)

	log.Fatal(http.ListenAndServe(":8315", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, insults.GenerateInsult())
}
