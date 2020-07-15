package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()
	for {
		con, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go serve(con)
	}
}

func serve(con net.Conn) {
	defer con.Close()
	var i int
	var method, url string
	scanner := bufio.NewScanner(con)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			xs := strings.Fields(ln)
			method = xs[0]
			url = xs[1]
			fmt.Println("METHODS: ", method)
			fmt.Println("URL: ", url)
		}
		if ln == "" {
			fmt.Println("THE END OF THE HTTP REQUEST HEADER")
			break
		}
		i++
	}
	body := `
		<!DOCTYPE html>
			<head>
				<meta charset="UTF-8">
				<title>hello</title>
			</head>
			<body>
				<h1>how are you</h1>
			</body>
		</html>
	`
	io.WriteString(con, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(con, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(con, "Content-Type: text/html\r\n")
	io.WriteString(con, "\r\n")
	io.WriteString(con, body)
}
