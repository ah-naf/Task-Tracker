package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
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

func CreateOpenFile(path string) (*os.File, error) {
	// Get the full path of the currently executing file (executable)
	executable, err := os.Executable()
	if err != nil {
		log.Fatalln("Error executing executable:", err)
	}

	// Get the directory where the executable is located
	dir := filepath.Dir(executable)
	
	file, err := os.OpenFile(filepath.Join(dir, path), os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalln("Error creating or opening file", err)
	}
	return file, nil
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

func ReadTaskFromFile(file *os.File) []Task {
	scanner := bufio.NewScanner(file)
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	if err := scanner.Err(); err!= nil {
        log.Fatalln("Error reading file:", err)
    }

	var task []Task
	if err := json.Unmarshal(buf.Bytes(), &task); err != nil {
		// log.Println("Error unmarshalling JSON:", err)
		return []Task{}
	}

	return task
}

func WriteTaskToFile(file *os.File, tasks []Task) {
	// Reset the file offset to the beginning for writing
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		log.Fatalln("Error seeking file:", err)
	}

	// Clear the file content before writing
	if err := file.Truncate(0); err != nil {
		log.Fatalln("Error truncating file:", err)
	}

	// Create a JSON encoder that writes directly to the file
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")

	// Encode and write the tasks array to the file
	if err := encoder.Encode(tasks); err != nil {
		log.Fatalln("Error encoding tasks to JSON:", err)
	}
}