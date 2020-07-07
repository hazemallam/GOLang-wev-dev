package main

import (
	"strings"
	"io"
	"log"
	"os"
	"fmt"
)
func main(){
	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	str := fmt.Sprint(
		`<!DOCTYPE html>
		<head>
		<meta charset="UTF-8">
		<title>Hello world</title>
		</head>
		<body>
		<h1>`+ name +`</h1>
		</body>
		</html>
	`)
	fileCreated, err := os.Create("index.html")
	if err != nil{
		log.Fatalln("error creating file", fileCreated)
	}
	io.Copy(fileCreated, strings.NewReader(str))
}