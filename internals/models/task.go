package models

import "time"

type Task struct {
	Id          int 		`json:"id"`
	Description string  	`json:"description"`
	Status      string  	`json:"status"`
	CreatedAt   time.Time 	`json:"createdAt"`
	UpdatedAt  	time.Time 	`json:"updatedAt"`
}

func NewTask(id int, description string, status string, createdAt time.Time, updatedAt time.Time) *Task {
	return &Task{Id: id, Description: description, Status: status, CreatedAt: createdAt, UpdatedAt: updatedAt}
}