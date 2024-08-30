package main

import (
	"log"
	"os"
	"path/filepath"
)

func CreateOpenFile(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalln("Error creating or opening file", err)
	}
	return file, nil
}

func handleCommand(args []string) {
	command := args[0]
	switch command {
	case "add":
	case "update":
	case "delete":
	case "list":
	case "start":
	case "complete":
	default:
		log.Fatalf("Error: Unknown command '%s'. Supported commands: add, update, delete, list, start, complete", command)
	}
}

func main()  {
	// Get the full path of the currently executing file (executable)
	executable, err := os.Executable()
	if err != nil {
		log.Fatalln("Error executing executable:", err)
	}

	// Get the directory where the executable is located
	dir := filepath.Dir(executable)

	// Create data.json file to store the task
	jsonFile, _ := CreateOpenFile(filepath.Join(dir, "data.json"))
	defer jsonFile.Close()

	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal(`Error: Missing command. Please provide a command (e.g., add, update, delete, mark-in-progress, mark-done, list).
Usage: task-tracker <command> [options]`)
	}
	
	handleCommand(args)
}