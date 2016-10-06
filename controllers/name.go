package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// NameController is the controller of the root path of the API
type NameController struct {
}

func (n *NameController) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name := vars["name"]

	fmt.Fprintf(w, "Hello %v", name)
}
