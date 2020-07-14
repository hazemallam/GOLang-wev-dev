package main

import (
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
		go handle(con)
	}
}

func handle(con net.Conn) {
	defer con.Close()
	io.WriteString(con, "i see your connection")

}
