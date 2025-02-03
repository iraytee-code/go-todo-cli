package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

// define a status type
type Status int

const (
	Completed Status = iota + 1
	NotCompleted
	InProgress
)

func (w Status) String() string {
	return [...]string{"Completed", "Not Completed", "In Progress"}[w-1]
}

func (w Status) EnumIndex() int {
	return int(w)
}

type item struct {
	ID          int `json:"id"`
	Task        string
	Status      Status
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []item



func SaveToFile(v *Todos) {
	CreateFile()
	bytes, _ := json.Marshal(v)
	err := os.WriteFile(
		filePath, bytes, 0644,
	)
	if err != nil {
		log.Fatal("Cannot save file", err)
	}
}

func (t *Todos) Add(task string) Todos {
	retrieved := ReadFile()
	id := len(retrieved) + 1

	todo := item{
		ID:        id,
		Task:      task,
		Status:    NotCompleted,
		CreatedAt: time.Now(),
	}
	*t = append(*t, todo)
	retrieved = append(retrieved, todo)
	SaveToFile(&retrieved)
	return *t
}
