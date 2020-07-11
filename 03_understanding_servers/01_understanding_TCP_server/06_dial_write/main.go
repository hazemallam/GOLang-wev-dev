package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	con, err := net.Dial("tcp", "localhost:5050")
	if err != nil {
		log.Fatalln(err)
	}
	defer con.Close()

	fmt.Fprintf(con, "hello read this form the dial writer")
}
