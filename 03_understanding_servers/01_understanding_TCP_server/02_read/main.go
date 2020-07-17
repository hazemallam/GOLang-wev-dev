package main

import (
	"bufio"
	"fmt"
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
	scanner := bufio.NewScanner(con)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}

	// we never get here
	// we have an open stream connection
	// how does the above reader know when it's done?
	fmt.Println("Code got here.")
}
