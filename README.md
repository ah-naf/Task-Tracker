# Task Tracker CLI

A simple command-line application to manage your tasks. 

[This Project is from roadmap.sh](https://roadmap.sh/projects/task-tracker)

## Features

- Add a new task with a title and an optional description.
- Update an existing task by ID with a new title or description.
- Delete a task by ID.
- List all tasks or filter tasks by their status (e.g., `done`, `in-progress`, `not-started`).
- Change the status of a task by ID.

## Usage

### General Usage

```bash
task-tracker <command> [options]
```

### Commands

1. **Add a Task**

   ```bash
   task-tracker add "Task Title" "Optional Description"
   ```

   - Adds a new task with the specified title and description.
   - The description is optional.

2. **Update a Task**

   ```bash
   task-tracker update <Task ID> "New Title" -desc "New Description"
   ```

   - Updates the title and/or description of a task by ID.
   - If the `-desc` flag is used, it updates only the description.

3. **Delete a Task**

   ```bash
   task-tracker delete <Task ID>
   ```

   - Deletes the task with the specified ID.

4. **List Tasks**

   ```bash
   task-tracker list
   ```

   - Lists all tasks.

   ```bash
   task-tracker list <status>
   ```

   - Lists tasks filtered by the specified status (`done`, `in-progress`, `not-started`).

5. **Change Task Status**

   ```bash
   task-tracker change-status <Task ID> <new status>
   ```

   - Changes the status of the task with the specified ID to the new status (`done`, `in-progress`, `not-started`).

## Example

```bash
task-tracker add "Buy groceries" "Remember to buy milk"
task-tracker update 1 "Buy groceries and cook dinner"
task-tracker change-status 1 in-progress
task-tracker list
task-tracker delete 1
```

## Notes

- Tasks are stored in a `data.json` file in the same directory as the executable.
- Task IDs are positive integers, starting from 1.
- The application automatically manages task IDs to ensure no gaps after a deletion.

## License

This project is licensed under the MIT License.
