package main

import (
	"log"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func HandleAdd(args []string) {
	if len(args) == 0 {
		log.Fatalln("Error: Missing title for the task. Title is required, description is optional.")
	}
	title, desc := args[0], ""

	if len(args) > 1 {
		desc = args[1]
	}

	file, _ := CreateOpenFile("data.json")
	defer file.Close()

	// Read Task from file
	tasks := ReadTaskFromFile(file)
	
	newTask := Task{
		ID:          len(tasks) + 1,
		Title: 		 title,
        Description: desc,
        Status:      "in progress",
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)

	// Write task to file
	WriteTaskToFile(file, tasks)

}
