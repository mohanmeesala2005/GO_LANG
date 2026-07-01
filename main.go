package main

import(
	"fmt"
	"os"
	"strconv"
)

func main(){
	args := os.Args[1:]

	if len(args) == 0{
		printUsage()
		os.Exit(1)
	}

	command := args[0]

	switch command {
	case "add" :
		if len(args) < 2{
			fmt.Println("Error: Task Description is required")
			os.Exit(1)
		}
		err := addTask(args[1])
		if err != nil{
			fmt.Println("Error : %v\n", err)
			os.Exit(1)
		}
	case "update":
		if len(args) < 1{
			fmt.Println("Error: Task Id and new description required")
			os.Exit(1)
		}

		id, err := strconv.Atoi(args[1])

		if err != nil{
			fmt.Println("Error: Id must be a number")
		os.Exit(1)
		}
		id, err = strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error: ID must be a number")
			os.Exit(1)
		}
		err = updateTask(id, args[2])
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

	case "delete":
		if len(args) < 2 {
			fmt.Println("Error: Task ID required")
			os.Exit(1)
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error: ID must be a number")
			os.Exit(1)
		}
		err = deleteTask(id)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

	case "mark-in-progress":
		if len(args) < 2 {
			fmt.Println("Error: Task ID required")
			os.Exit(1)
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error: ID must be a number")
			os.Exit(1)
		}
		err = markTask(id, StatusInProgress)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

	case "mark-done":
		if len(args) < 2 {
			fmt.Println("Error: Task ID required")
			os.Exit(1)
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error: ID must be a number")
			os.Exit(1)
		}
		err = markTask(id, StatusDone)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

	case "list":
		filter := Status("")
		if len(args) >= 2 {
			switch args[1] {
			case "done":
				filter = StatusDone
			case "todo":
				filter = StatusPending
			case "in-progress":
				filter = StatusInProgress
			default:
				fmt.Printf("Unknown filter: %s\n", args[1])
				os.Exit(1)
			}
		}
		err := listTasks(filter)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func markTask(id int, status Status) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	idx, err := findTaskByID(tasks, id)
	if err != nil {
		return err
	}

	tasks[idx].Status = status

	return saveTasks(tasks)
}

func printUsage() {
	fmt.Println("Task Tracker CLI")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  task-cli add \"<description>\"")
	fmt.Println("  task-cli update <id> \"<description>\"")
	fmt.Println("  task-cli delete <id>")
	fmt.Println("  task-cli mark-in-progress <id>")
	fmt.Println("  task-cli mark-done <id>")
	fmt.Println("  task-cli list [done|todo|in-progress]")
}