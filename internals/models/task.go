package models

import (
	"fmt"
	"time"
)

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

func (task Task) Display() string {
	return "Task{" +
		"id=" + fmt.Sprint(task.Id) +
		", description=" + task.Description +
		", status=" + task.Status +
		", createdAt=" + task.CreatedAt.String() +
		", updatedAt=" + task.UpdatedAt.String() +
		"}"
}