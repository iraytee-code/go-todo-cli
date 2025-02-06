package main

import (
	"encoding/json"
	"fmt"
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

// Error implements error.
func (t Todos) Error() string {
	panic("unimplemented")
}

func (s *Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *Status) UnmarshalJSON(data []byte) error {
	var statusStr string
	if err := json.Unmarshal(data, &statusStr); err != nil {
		return err
	}
	switch statusStr {
	case "Completed":
		*s = Completed
	case "Not Completed":
		*s = NotCompleted
	case "In Progress":
		*s = InProgress
	default:
		return fmt.Errorf("invalid status: %s", statusStr)
	}

	return nil
}

func SaveToFile(v *Todos) error {
	CreateFile()
	bytes, _ := json.Marshal(v)
	err := os.WriteFile(
		filePath, bytes, 0644,
	)
	if err != nil {
		log.Fatal("Cannot save file", err)
	}
	return err
}

func (t *Todos) Add(task string) error {
	retrieved := ReadFile()
	// Find highest ID
	maxID := 0
	for _, item := range retrieved {
		if item.ID > maxID {
			maxID = item.ID
		}
	}

	todo := item{
		ID:        maxID + 1,
		Task:      task,
		Status:    NotCompleted,
		CreatedAt: time.Now(),
	}

	retrieved = append(retrieved, todo)
	*t = retrieved

	return SaveToFile(&retrieved)

}

func List() ([]item, error) {
	return ReadFile(), nil
}

func Delete(taskID int) ([]item, error) {
	currentTasks := ReadFile()
	var err error
	currentTasks, err = removeById(currentTasks, taskID)
	if err != nil {
		return nil, err
	}
	for i := range currentTasks {
		currentTasks[i].ID = i + 1
	}
	SaveToFile(&currentTasks)
	return currentTasks, nil
}

func removeById(tasks []item, targetId int) ([]item, error) {
	for i, task := range tasks {
		if task.ID == targetId {
			return append(tasks[:i], tasks[i+1:]...), nil
		}
	}
	return tasks, fmt.Errorf("task with ID %d not found", targetId)
}
