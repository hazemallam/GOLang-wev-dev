package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":5050", nil)

}

func foo(res http.ResponseWriter, req *http.Request) {
	value := req.FormValue("name")
	res.Header().Set("content-type", "text/html; charset=utf-8")
	io.WriteString(res, `
		<!DOCTYPE html>
			<head>
				<meta charset="UTF-8">
				<title>state</title>
			</head>
			<body>
				<form method="GET">
					<input type="text" name="name">
					<input type="submit">
				</form>
			</body>
		</html>
	`+value)
}
