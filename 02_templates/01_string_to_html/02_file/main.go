package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "James Bond"
	str := fmt.Sprint(
		`<!DOCTYPE html>
		<head>
		<meta charset="UTF-8">
		<title>Hello world</title>
		</head>
		<body>
		<h1>` + name + `</h1>
		</body>
		</html>
	`)
	fileCreated, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("error creating file", err)
	}
	defer fileCreated.Close()
	io.Copy(fileCreated, strings.NewReader(str))
}
