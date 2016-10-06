package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lfmexi/rest-api-example/controllers"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	router.Handle("/hello/{name}", &controllers.NameController{})

	log.Fatal(http.ListenAndServe(":8080", router))
}
