package main

import "fmt"

func main() {
	name := "James Bond"

	template := `
	<!DOCTYPE html>
	<head>
	<meta charset="UTF-8">
	<title>Hello world</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`
	fmt.Println(template)
}
