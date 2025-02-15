package main

import (
	"log"
	"os"
)

// handleCommand routes the provided command to the appropriate handler function
func handleCommand(args []string) {
	command := args[0]
	switch command {
	case "add":
		HandleAdd(args[1:])
	case "update":
		HandleUpdate(args[1:])
	case "delete":
		HandleDelete(args[1:])
	case "list":
		HandleList(args[1:])
	case "change-status":
		HandleChangeStauts(args[1:])
	default:
		log.Fatalf("Error: Unknown command '%s'. Supported commands: add, update, delete, list, change-status", command)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal(`Error: Missing command. Please provide a command (e.g., add, update, delete, list, change-status).
Usage: task-tracker <command> [options]`)
	}
	handleCommand(args)
}
