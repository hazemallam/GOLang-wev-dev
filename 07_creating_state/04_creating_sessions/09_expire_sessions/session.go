package main

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func getUser(res http.ResponseWriter, req *http.Request) user {
	//get cookie
	cookie, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewRandom()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	cookie.MaxAge = sessionLength
	http.SetCookie(res, cookie)

	//if user already exists, get user
	var u user
	if s, ok := dbSessions[cookie.Value]; ok {
		s.lastActivity = time.Now()
		dbSessions[cookie.Value] = s
		u = dbUsers[s.un]
	}
	return u
}

func loggedIn(res http.ResponseWriter, req *http.Request) bool {
	cookie, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := dbSessions[cookie.Value]
	if ok {
		s.lastActivity = time.Now()
		dbSessions[cookie.Value] = s
	}

	_, ok = dbUsers[s.un]

	//refresh session
	cookie.MaxAge = sessionLength
	http.SetCookie(res, cookie)
	return ok
}

func cleanSessions() {
	fmt.Println("BEFOR CLEAN")
	showMessage()
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN")
	showMessage()
}

//for demonstration purposes
func showMessage() {
	fmt.Println("*******************")
	for k, v := range dbSessions {
		fmt.Println(k, v.un)
	}
	fmt.Println("")
}
