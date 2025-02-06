package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var filePath = "./todo.json"

func CreateFile() error {
	todo := []Todos{}
	jsonData, err := json.MarshalIndent(todo, "", "  ")
	if err != nil {
		log.Fatal("Cannot marshal json", err)
	}
	err = os.WriteFile(
		filePath, jsonData, 0644,
	)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	return err
}

func ReadFile() Todos {
	CheckFileExists()
	var allTodo Todos
	todos, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Cannot read file", err)
	}
	err = json.Unmarshal(todos, &allTodo)
	if err != nil {
		log.Fatal("Cannot retrieve json", err)
	}
	return allTodo
}

func CheckFileExists() error {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Println("File does not exist, Creating a new file..")
		CreateFile()
	}
	return err
}
