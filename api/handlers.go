package api

import (
	"encoding/json"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	"github.com/rhass99/pharmacy-locum/storage"
	//"github.com/satori/go.uuid"
	//"fmt"
	"log"
	"net/http"
	"text/template"
)

type User struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

type Session struct {
	Id              string
	Authenticated   bool
	Unauthenticated bool
	User            User
}

var signupTmpl = template.Must(template.New("signup.html").ParseFiles("tmpl/signup.html"))
var loginTmpl = template.Must(template.New("login.html").ParseFiles("tmpl/login.html"))
var profileTmpl = template.Must(template.New("profile.html").ParseFiles("tmpl/profile.html"))

var store = sessions.NewCookieStore([]byte(storage.RandId(50)))

var db storage.Store

func ProfileApplicantGET(w http.ResponseWriter, r *http.Request) {
	err := profileTmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func LoginApplicantGET(w http.ResponseWriter, r *http.Request) {
	err := loginTmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func LoginApplicantGET(w http.ResponseWriter, r *http.Request) {
	err := loginTmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func SignUpApplicantGET(w http.ResponseWriter, r *http.Request) {
	err := signupTmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func SignUpApplicantPOST(w http.ResponseWriter, r *http.Request) {
	var a storage.Applicant
	db.Path = "/Users/rami/go/src/github.com/rhass99/pharmacy-locum/db/applicants.db"
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	decoder := schema.NewDecoder()
	err = decoder.Decode(&a, r.PostForm)
	if err != nil {
		log.Println(err)
	}

	a.ID = storage.RandId(50)
	a.Password = string(storage.WeakPasswordHash(a.Password))

	err = db.Open("Applicant")
	if err != nil {
		log.Println(err)
	}

	err = db.CreateApplicant(&a)
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
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}
