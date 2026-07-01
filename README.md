# Task Tracker CLI

A lightweight, zero-dependency command-line task manager written in Go. Tasks are persisted locally to a JSON file — no database or internet connection required.

> **GitHub Repository:** [https://github.com/mohanmeesala2005/GO_LANG](https://github.com/mohanmeesala2005/GO_LANG)

---

## Features

### ➕ Add a Task
Create a new task with a description. Each task is automatically assigned a unique incremental ID and given a default status of **In Progress**.

```bash
task-cli add "Buy groceries"
```

---

### ✏️ Update a Task
Change the description of an existing task by its ID.

```bash
task-cli update 1 "Buy groceries and cook dinner"
```

---

### 🗑️ Delete a Task
Permanently remove a task by its ID.

```bash
task-cli delete 1
```

---

### 🔄 Mark a Task as In Progress
Set the status of a task to **In Progress**.

```bash
task-cli mark-in-progress 2
```

---

### ✅ Mark a Task as Done
Set the status of a task to **Done**.

```bash
task-cli mark-done 2
```

---

### 📋 List Tasks
Display all tasks in a formatted table showing ID, Status, Description, and last-updated timestamp.

```bash
task-cli list
```

You can also **filter** the list by status:

| Filter Flag    | Shows                    |
|----------------|--------------------------|
| `done`         | Only completed tasks      |
| `todo`         | Only pending tasks        |
| `in-progress`  | Only in-progress tasks    |

```bash
task-cli list done
task-cli list todo
task-cli list in-progress
```

---

## Task Data Model

Every task stores the following fields:

| Field         | Type     | Description                              |
|---------------|----------|------------------------------------------|
| `id`          | int      | Auto-incremented unique identifier       |
| `description` | string   | Text describing the task                 |
| `status`      | string   | One of: `pending`, `Inprogress`, `Done`  |
| `createdAt`   | time     | Timestamp when the task was created      |
| `updatedAt`   | time     | Timestamp of the last modification       |

---

## Data Storage

Tasks are saved to a `tasks.json` file in the same directory where you run the CLI. The file is created automatically on first use — no setup needed.

---

## Getting Started

### Prerequisites
- [Go 1.21+](https://golang.org/dl/)

### Build & Run

```bash
# Clone the repository
git clone https://github.com/mohanmeesala2005/GO_LANG
cd task_cli

# Build the binary
go build -o task-cli .

# Run it
./task-cli add "My first task"
./task-cli list
```

---

## Usage Reference

```
task-cli add "<description>"
task-cli update <id> "<new description>"
task-cli delete <id>
task-cli mark-in-progress <id>
task-cli mark-done <id>
task-cli list [done|todo|in-progress]
```

---

## Project Structure

```
task_cli/
├── main.go       # Entry point & CLI argument routing
├── commands.go   # Task command implementations (add, update, delete, list)
├── file.go       # JSON file persistence (load, save, helpers)
├── task.go       # Task struct and Status type definitions
├── tasks.json    # Auto-generated local task storage
└── go.mod        # Go module definition
```