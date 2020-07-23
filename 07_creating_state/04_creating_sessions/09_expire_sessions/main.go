package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}
type session struct {
	un           string
	lastActivity time.Time
}

var tpl *template.Template
var dbUsers = map[string]user{}       //user ID, user
var dbSessions = map[string]session{} //session ID, session
var dbSessionsCleaned time.Time

const sessionLength = 30

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
	dbSessionsCleaned = time.Now()
}

func main() {

}

func index(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	showMessage()
	tpl.ExecuteTemplate(res, "index.html", u)
}

func bar(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	if !loggedIn(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "admin" {
		http.Error(res, "YOU ARE NOT ALLOWED TO ENTER THIS PAGE", http.StatusForbidden)
		return
	}
	showMessage()
	tpl.ExecuteTemplate(res, "bar.html", u)
}

func signup(res http.ResponseWriter, req *http.Request) {
	if loggedIn(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	var u user
	//process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("uname")
		p := req.FormValue("password")
		f := req.FormValue("fname")
		l := req.FormValue("lname")
		r := req.FormValue("role")

		//username taken?
		if _, ok := dbUsers[un]; ok {
			http.Error(res, "username already taken", http.StatusForbidden)
			return
		}

		//creating session
		sID, _ := uuid.NewRandom()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		cookie.MaxAge = sessionLength
		http.SetCookie(res, cookie)
		dbSessions[cookie.Value] = session{un, time.Now()}

		//store user in db
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(res, "Internal server error", http.StatusInternalServerError)
			return
		}
		u = user{un, bs, f, l, r}
		dbUsers[un] = u

		//redirect
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	showMessage()
	tpl.ExecuteTemplate(res, "signup.html", u)
}

func login(res http.ResponseWriter, req *http.Request) {
	if loggedIn(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	var u user

	//from submession
	if req.Method == http.MethodPost {
		un := req.FormValue("uname")
		p := req.FormValue("password")

		//check username
		u, ok := dbUsers[un]
		if !ok {
			http.Error(res, "username or password incorrect", http.StatusForbidden)
			return
		}

		//check password
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(res, "username or password incorrect", http.StatusForbidden)
			return
		}

		//create session
		sID, _ := uuid.NewRandom()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		cookie.MaxAge = sessionLength
		http.SetCookie(res, cookie)
		dbSessions[cookie.Value] = session{un, time.Now()}
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	showMessage()
	tpl.ExecuteTemplate(res, "login.html", u)
}

func logout(res http.ResponseWriter, req *http.Request) {
	if !loggedIn(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	cookie, _ := req.Cookie("session")

	//delete session
	delete(dbSessions, cookie.Value)

	//remove the cookie
	cookie = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(res, cookie)

	//clean up dbsessions
	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
		go cleanSessions()
	}
	http.Redirect(res, req, "/login", http.StatusSeeOther)
	return
}
