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
		go handle(con)
	}
}

func handle(con net.Conn) {
	defer con.Close()
	scanner := bufio.NewScanner(con)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}

	fmt.Println("code got here")
	io.WriteString(con, "I see you connected")
}
