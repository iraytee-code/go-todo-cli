package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting...")
	var fileExists = CheckFileExists()
	if fileExists != nil {
		CreateFile()
	}
	var todos Todos
	todos.Add("Study at work...")
}
