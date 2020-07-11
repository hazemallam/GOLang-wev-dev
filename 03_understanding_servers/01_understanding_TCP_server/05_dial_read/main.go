package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	con, err := net.Dial("tcp", "localhost:5050")
	if err != nil {
		log.Fatalln(err)
	}
	defer con.Close()

	bs, err := ioutil.ReadAll(con)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))
}
