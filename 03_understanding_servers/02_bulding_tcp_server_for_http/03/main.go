package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		con, err := li.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		go handle(con)
	}
}

func handle(con net.Conn) {
	defer con.Close()
	request(con)
}

func request(con net.Conn) {
	i := 0
	scanner := bufio.NewScanner(con)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(con, ln)
		}
		if ln == "" {
			break
		}
		i++
	}
}

func mux(con net.Conn, ln string) {
	//request line
	m := strings.Fields(ln)[0]   //method
	uri := strings.Fields(ln)[1] //uri
	fmt.Println("***METHOD", m)
	fmt.Println("***URI", uri)

	if m == "GET" && uri == "/" {
		index(con)
	}
	if m == "GET" && uri == "/about" {
		about(con)
	}
	if m == "GET" && uri == "/contact" {
		contact(con)
	}
	if m == "GET" && uri == "/apply" {
		apply(con)
	}
	if m == "POST" && uri == "/apply" {
		applyProcess(con)
	}

}

func index(con net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>INDEX</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`
	fmt.Fprint(con, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(con, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(con, "Content-Type: text/html\r\n")
	fmt.Fprint(con, "\r\n")
	fmt.Fprint(con, body)
}

func about(con net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>ABOUT</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`

	fmt.Fprint(con, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(con, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(con, "Content-Type: text/html\r\n")
	fmt.Fprint(con, "\r\n")
	fmt.Fprint(con, body)
}

func contact(con net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>CONTACT</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`

	fmt.Fprint(con, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(con, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(con, "Content-Type: text/html\r\n")
	fmt.Fprint(con, "\r\n")
	fmt.Fprint(con, body)
}

func apply(con net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body>
	<strong>APPLY</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	<form method="POST" action="/apply">
	<input type="submit" value="apply">
	</form>
	</body></html>`

	fmt.Fprint(con, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(con, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(con, "Content-Type: text/html\r\n")
	fmt.Fprint(con, "\r\n")
	fmt.Fprint(con, body)
}

func applyProcess(con net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>APPLY PROCESS</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`

	fmt.Fprint(con, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(con, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(con, "Content-Type: text/html\r\n")
	fmt.Fprint(con, "\r\n")
	fmt.Fprint(con, body)
}
