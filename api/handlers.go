package api

import (
	"encoding/json"
	"github.com/rhass99/pharmacy-locum/storage"
	"log"
	"net/http"
)

var db storage.Store

func SignApplicant(w http.ResponseWriter, r *http.Request) {
	var a *storage.Applicant
	db.Path = "/Users/rami/go/src/github.com/rhass99/pharmacy-locum/db/applicants.db"
	defer db.Close()

	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		log.Println(err)
	}

	a.ID = storage.RandId(20)

	err = db.Open("Applicant")
	if err != nil {
		log.Println(err)
	}

	err = db.CreateApplicant(a)
	if err != nil {
		log.Println(err)
	}

	aback, err := db.GetApplicants()
	if err != nil {
		log.Println(err)
	}
	j, err := json.Marshal(aback)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}
