package main   

import (
	"encoding/json"
	"errors"
	"os"
)

const fileName = "tasks.json"

func loadTasks() ([]Task, error){
	if _,err := os.Stat(fileName); os.IsNotExist(err){
		return []Task{}, nil
	}

	data,err := os.ReadFile(fileName)
	if err != nil{
		return nil, err
	}

	if len(data) == 0{
		return []Task{}, nil
	}

	var tasks []Task
	err = json.Unmarshal(data,&tasks)

	if err != nil{
		return nil, err
	}

	return tasks, nil
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")

	if err != nil{
		return err
	}

	err = os.WriteFile(fileName, data, 0644)
	if err != nil{
		return err
	}

	return nil
}

func getNextId(tasks []Task) int{
	maxId := 0

	for _, task := range tasks{
		if task.Id > maxId {
			maxId = task.Id
		}
	}

	return  maxId + 1
}

func findTaskByID(tasks []Task, id int) (int ,error){
	for i, task := range tasks{
		if task.Id == id{
			return i, nil
		}
	}

	return -1, errors.New("Task not found")
}