package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/olekukonko/tablewriter"
)

func DisplayTodos(todos Todos) {
    if len(todos) == 0 {
        fmt.Println("No todos found.")
        return
    }

    table := tablewriter.NewWriter(os.Stdout)
    table.SetHeader([]string{"ID", "Task", "Status", "Created At"})
    
    for _, todo := range todos {
        table.Append([]string{
            strconv.Itoa(todo.ID),
            todo.Task,
            todo.Status.String(),
            todo.CreatedAt.Format("2006-01-02 15:04:05"),
        })
    }

    table.Render()
}