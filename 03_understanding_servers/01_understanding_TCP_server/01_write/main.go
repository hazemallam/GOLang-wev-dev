package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main(){
	li, err := net.Listen("tcp", ":5050")
	if err != nil{
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		con, err := li.Accept()
		if err != nil{
			log.Println(err)
			continue
		}
		io.WriteString(con, "\nHello from tcp server")
		fmt.Fprintln(con, "How is your day?")
		fmt.Fprintf(con, "%v", "Well, I hope!")
		con.Close()

	}
}