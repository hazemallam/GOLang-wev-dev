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
		log.Fatalln(err)
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

	//read request
	request(con)

	//write request
	response(con)
}

func request(con net.Conn) {
	i := 0
	scanner := bufio.NewScanner(con)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			//request line
			m := strings.Fields(ln)[0]
			fmt.Println("***Methods", m)
		}
		if ln == "" {
			//headers are done
			break
		}
		i++
	}
}

func response(con net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`
	fmt.Fprint(con, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(con, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(con, "Content-Type: text/html\r\n")
	fmt.Fprint(con, "\r\n")
	fmt.Fprint(con, body)
}
