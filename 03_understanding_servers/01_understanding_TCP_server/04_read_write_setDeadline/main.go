package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
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
	err := con.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Fatalln("Connection timeout")
	}

	scanner := bufio.NewScanner(con)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(con, "I heard you say: %s\n", ln)
	}

	defer con.Close()

	// now we get here
	// the connection will time out
	// that breaks us out of the scanner loop
	fmt.Println("***CODE GOT HERE***")

}
