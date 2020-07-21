package main

import (
	"fmt"
	"github.com/google/uuid"
	"html/template"
	"net/http"
)

type user struct {
	UserName, First, Last string
}

var tpl *template.Template
var dbUsers = map[string]user{}      //user ID, user
var dbSessions = map[string]string{} //session ID, user ID

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":5050", nil)

}

func index(res http.ResponseWriter, req *http.Request) {
	//get cookie
	cookie, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewRandom()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	http.SetCookie(res, cookie)

	//if user already exists, get user
	var u user
	if un, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[un]
		fmt.Println("u value", u)

	}

	//process from submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		fname := req.FormValue("fname")
		lname := req.FormValue("lname")
		u = user{un, fname, lname}
		dbSessions[cookie.Value] = un
		dbUsers[un] = u
	}
	tpl.ExecuteTemplate(res, "index.html", u)
}

func bar(res http.ResponseWriter, req *http.Request) {
	//get cookie
	cookie, err := req.Cookie("session")
	if err != nil {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[cookie.Value]
	if !ok {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(res, "bar.html", u)
}
