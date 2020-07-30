package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/db")
	check(err)
	defer db.Close()
	insert, err := db.Query("INSERT INTO student VALUES ('20','HASSAN')")
	check(err)
	defer insert.Close()
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":5050", nil)
	check(err)
}

func index(res http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(res, "successfully completed")
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
