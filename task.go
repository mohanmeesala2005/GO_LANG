package main

import (
	"time"
)

type Status string

const(
	StatusPending Status = "pending"
	StatusDone Status = "Done"
	StatusInProgress Status = "Inprogress"
)

type Task struct {
	Id 				int 		`json:"id"`
	Description 	string 		`json:"description"`
	Status 			Status 		`json:"status"`
	CreatedAt 		time.Time   `json:"createdAt"`
	UpdatedAt 		time.Time 	`json:"updatedAt"`
}