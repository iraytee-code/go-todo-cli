package main

import (
	// "fmt"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: todo [command] [arguments]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCommand := flag.NewFlagSet("add", flag.ExitOnError)
		addCommand.Parse(os.Args[2:])
		task := strings.Join(addCommand.Args(), " ")
		if task == "" {
			fmt.Println("Error: empty todo item is not allowed")
			os.Exit(1)
		}
		allTodo := ReadFile()
		err := allTodo.Add(task)
		if err != nil {
			fmt.Println("Error saving todo:", err)
			os.Exit(1)
		}

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: please provide a todo ID to delete")
			os.Exit(1)
		}
		_, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: invalid ID format")
			os.Exit(1)
		}

	case "list":
		allTodo := ReadFile()
		DisplayTodos(
			allTodo,
		)

	default:
		fmt.Println("Unknown command. Use 'todo add', 'todo delete', or 'todo list'")
		os.Exit(1)
	}
}
