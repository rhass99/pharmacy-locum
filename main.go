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
	r.HandleFunc("/signup", api.SignUpApplicantPOST).Methods("POST")
	r.HandleFunc("/signup", api.SignUpApplicantGET).Methods("GET")
	r.HandleFunc("/login", api.LoginApplicantGET).Methods("GET")
	r.HandleFunc("/profile", api.ProfileApplicantGET).Methods("GET")
	r.HandleFunc("/login", api.LoginApplicantPOST).Methods("Post")

	http.ListenAndServe(":8080", r)
}
