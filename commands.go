package main

import (
	"fmt"
	"time"
)

func addTask(description string) error {
	tasks, err := loadTasks()
	if err != nil{
		return err
	}

	newTask := Task{
		Id : getNextId(tasks),
		Description: description,
		Status: StatusInProgress,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tasks = append(tasks, newTask)

	err = saveTasks(tasks)

	if err != nil{
		return err
	}

	fmt.Println("Task Added Successfully (Id : %d)\n", newTask.Id)

	return nil
}

func updateTask (id int, description string) error {
	tasks, err := loadTasks()
	if err != nil{
		return err
	}

	idx, err := findTaskByID(tasks,id)

	if err != nil{
		return err
	}

	tasks[idx].Description = description
	tasks[idx].UpdatedAt = time.Now()

	err = saveTasks(tasks)

	if err != nil{
		return err
	}

	fmt.Println("Task Updated Successfully")
	return nil
}

func deleteTask(id int) error{
	tasks, err := loadTasks()

	if err != nil{
		return err
	}

	idx, err := findTaskByID(tasks, id)

	if err != nil{
		return err
	}

	tasks = append(tasks[:idx], tasks[idx+1:]...)

	err = saveTasks(tasks)
	
	if err != nil{
		return err
	}

	fmt.Println("Task deleted!")

	return nil
}

func listTasks(filter Status) error{
	tasks, err := loadTasks()

	if err != nil{
		return err
	}

	if len(tasks) == 0{
		fmt.Println("No task found")
		return nil
	}

	fmt.Printf("%-5s %-15s %-30s %-20s\n", "ID", "Status", "Description", "Updated")
	fmt.Println(string(make([]byte, 75)))

	for _, task := range tasks {
		if filter != "" && task.Status != filter{
			continue;
		} 

		desc := task.Description

		if len(desc) > 28{
			desc = desc[:25] + "..."
		}

		fmt.Printf("%-5d %-15s %-30s %-20s\n",
			task.Id,
			task.Status,
			desc,
			task.UpdatedAt.Format("2006-01-02 15:04"),
		)
	}

	return nil
}