package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rhass99/pharmacy-locum/api"
	"net/http"
)

func handlerz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/handlerz", handlerz)
	r.HandleFunc("/signup", api.SignUpApplicant).Methods("POST")

	http.ListenAndServe(":8080", r)
}
