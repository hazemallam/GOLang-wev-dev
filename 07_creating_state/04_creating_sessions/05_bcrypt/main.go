package main

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
}

var tpl *template.Template
var dbSessions = map[string]string{} //session ID, user ID
var dbUsers = map[string]user{}      //user ID, user

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":5050", nil)

}

func index(res http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	tpl.ExecuteTemplate(res, "index.html", u)
}

func bar(res http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	fmt.Println("user: ", u)

	if !loggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "bar.html", u)
}

func signup(res http.ResponseWriter, req *http.Request) {
	if loggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	//process form submisson
	if req.Method == http.MethodPost {
		un := req.FormValue("uname")
		p := req.FormValue("password")
		f := req.FormValue("fname")
		l := req.FormValue("lname")

		//username taken ?
		if _, ok := dbUsers[un]; ok {
			http.Error(res, "username already taken", http.StatusForbidden)
			http.Redirect(res, req, "/", http.StatusSeeOther)
			return
		}

		//create session
		sID, _ := uuid.NewRandom()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(res, cookie)
		dbSessions[cookie.Value] = un

		//store user
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(res, "internal server error", http.StatusInternalServerError)
		}
		u := user{un, bs, f, l}
		dbUsers[un] = u

		//redirect
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(res, "signup.html", nil)
}
