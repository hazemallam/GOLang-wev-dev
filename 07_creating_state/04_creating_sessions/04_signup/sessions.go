package main

import (
	"fmt"
	"net/http"
)

func getUser(req *http.Request) user {
	var u user
	//get cookie
	cookie, err := req.Cookie("session")
	if err != nil {
		return u
	}
	if un, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[un]
		fmt.Println("user_1", u)

	}
	return u
}

func loggedIn(req *http.Request) bool {
	cookie, err := req.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSessions[cookie.Value]
	_, ok := dbUsers[un]
	return ok
}
