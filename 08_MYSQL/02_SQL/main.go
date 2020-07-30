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

	db, err = sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/godb")
	check(err)
	defer db.Close()
	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/amigos", students)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", del)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":5050", nil)
	check(err)

}

func index(res http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(res, "at index")
	check(err)
}

func students(res http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT name From student;`)
	check(err)
	var s, name string
	s = "RETREIVED RECORDS\n"

	//query
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(res, s)
}

func create(res http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`CREATE TABLE teacher (name VARCHAR(20));`)
	check(err)
	defer stmt.Close()

	result, err := stmt.Exec()
	check(err)

	n, err := result.RowsAffected()
	check(err)

	fmt.Fprintln(res, "CREATED TABLE teacher", n)
}

func insert(res http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO teacher VALUES ('James');`)
	check(err)
	defer stmt.Close()

	result, err := stmt.Exec()
	check(err)
	n, err := result.RowsAffected()
	check(err)

	fmt.Fprintln(res, "INSERTED RECORD", n)
}

func read(res http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT * FROM teacher;`)
	check(err)
	defer rows.Close()

	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		fmt.Fprintln(res, "RETREIVED RECORDS:", name)
	}
}

func update(res http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`UPDATE teacher SET name="Bond" WHERE name="James";`)
	check(err)
	defer stmt.Close()
	result, err := stmt.Exec()
	check(err)
	n, err := result.RowsAffected()
	check(err)
	fmt.Fprintln(res, "UPDATED RECORD", n)
}

func del(res http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DELETE FROM teacher WHERE name="Bond";`)
	check(err)
	defer stmt.Close()

	result, err := stmt.Exec()
	check(err)

	n, err := result.RowsAffected()
	check(err)

	fmt.Fprintln(res, "DELETED RECORED", n)
}

func drop(res http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DROP TABLE teacher;`)
	check(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	check(err)

	fmt.Fprintln(res, "DROPPED TABLE teacher")
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
