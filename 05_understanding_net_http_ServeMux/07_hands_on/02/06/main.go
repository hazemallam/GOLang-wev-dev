package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
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
	scanner := bufio.NewScanner(con)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			fmt.Println("THIS IS THE END OF HTTP REQUEST HEADER")
			break
		}
	}
	body := "CHECK OUT THE RESPONSE BODY"
	io.WriteString(con, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(con, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(con, "Content-Type: text/plain\r\n")
	io.WriteString(con, "\r\n")
	io.WriteString(con, body)

}
