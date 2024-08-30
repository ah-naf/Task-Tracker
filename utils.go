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
        Status:      "not-started",
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

func HandleDelete(args []string) {
	if len(args) < 1 {
        log.Fatalf("Error: Missing ID for the task to be deleted. ID is required.")
    }
    idStr := args[0]

    id, err := strconv.Atoi(idStr)
    if err!= nil || id <= 0 {
        log.Fatalf("Error: Invalid ID '%s'. ID should be a positive integer.", idStr)
    }

    file, _ := CreateOpenFile("data.json")
    defer file.Close()

    tasks := ReadTaskFromFile(file)

    deleted := false
	tempID := id
    for i, task := range tasks {
        if task.ID == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            deleted = true
        } else if deleted {
			tasks[i-1].ID = tempID
			tempID++
		}
    }

    if!deleted {
        log.Fatalf("Error: No task found with ID %d.", id)
    }

    WriteTaskToFile(file, tasks)
	log.Println("Task deleted successfully.")
}

func HandleList(args []string) {
	file, _ := CreateOpenFile("data.json")
    defer file.Close()
	tasks := ReadTaskFromFile(file)

	if len(args) == 0 {
		// List all tasks
        for _, task := range tasks {
            PrintSingleTask(task)
        }
        return
	} else {
		option := args[0]
		// option can be done, in progress, not-started
		switch option {
        case "done":
            for _, task := range tasks {
                if task.Status == "done" {
                    PrintSingleTask(task)
                }
            }
            return
        case "in-progress":
            for _, task := range tasks {
                if task.Status == "in-progress" {
					PrintSingleTask(task)
				}
			}
		case "not-started":
            for _, task := range tasks {
                if task.Status == "not-started" {
                    PrintSingleTask(task)
                }
            }
		default:
			log.Fatalf("Error: Invalid option '%s'. Supported options: done, in-progress, not-started.", option)
		}
	}
}

func HandleChangeStauts(args []string) {
	if len(args) < 2 {
        log.Fatalf("Error: Missing ID and Status for the task to be changed. ID and Status is required.")
    }
	idStr := args[0]
	changedStatus := args[1]
	if changedStatus!= "done" && changedStatus!= "in-progress" && changedStatus!= "not-started" {
        log.Fatalf("Error: Invalid status '%s'. Supported statuses: done, in-progress, not-started.", changedStatus)
    }

	id, err := strconv.Atoi(idStr)
    if err!= nil || id <= 0 {
        log.Fatalf("Error: Invalid ID '%s'. ID should be a positive integer.", idStr)
    }

	file, _ := CreateOpenFile("data.json")
    defer file.Close()

    tasks := ReadTaskFromFile(file)

	updated := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = changedStatus
			updated = true
			break
		}
	}
	if !updated {
		log.Fatalf("Error: No task found with ID %d.", id)
	}

	WriteTaskToFile(file, tasks)
	log.Println("Status updated successfully")
}

func PrintSingleTask(task Task) {
	log.Printf("ID: %d, Title: %s, Description: %s, Status: %s, Created At: %s, Updated At: %s\n",
		task.ID, task.Title, task.Description, task.Status, task.CreatedAt.Format(time.RFC3339), task.UpdatedAt.Format(time.RFC3339))
}