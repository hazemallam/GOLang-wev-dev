package main

import (
	"fmt"
	"io"
	"bufio"
	"net"
	"log"
)

func main(){
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

		scanner := bufio.NewScanner(con)
		for scanner.Scan(){
			ln := scanner.Text()
			fmt.Println(ln)
			if ln == "" {
				fmt.Println("THIS IS THE END OF THE HTTP CONNECTION")
				break
			}
		}
		fmt.Println("CODE GET HERE")
		io.WriteString(con, "I see you connected")
		con.Close()
	}
}