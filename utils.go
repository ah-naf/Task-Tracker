package main

import (
	"log"
	"strconv"
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

	log.Println("Task added successfully")

}

func HandleUpdate(args []string) {
	// Task update need an ID and Title. Description is optional.
	if len(args) < 2 {
        log.Fatalf("Error: Missing ID and/or Title for the task. ID and Title are required.")
    }
    idStr := args[0]

    id, err := strconv.Atoi(idStr)
    if err!= nil || id <= 0 {
        log.Fatalf("Error: Invalid ID '%s'. ID should be a positive integer.", idStr)
    }

	var title, desc string
	for i := 1; i < len(args); i++ {
		if args[i] == "-desc" {
			if i+1 < len(args) {
				desc = args[i+1]
				break
			} else {
				log.Fatalf("Error: Missing description after '-desc' flag.")
			}
		} else if title == "" {
			title = args[i]
		} else if title != "" && desc == "" {
			desc = args[i]
			break
		}
	}

	if title == "" && desc == "" {
		log.Fatalln("Error: Neither Title nor Description provided for the update. At least one must be specified.")
	}

    file, _ := CreateOpenFile("data.json")
    defer file.Close()

    tasks := ReadTaskFromFile(file)

	updated := false
    for i, task := range tasks {
        if task.ID == id {
            if title != "" {
				tasks[i].Title = title
			}
			if desc != "" {
				tasks[i].Description = desc
			}
			tasks[i].UpdatedAt = time.Now()
			updated = true
			break
		}
	}

	if !updated {
		log.Fatalf("Error: No task found with ID %d.", id)
	}

	WriteTaskToFile(file, tasks)

	log.Println("Task updated successfully.")
}
