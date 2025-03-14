package service

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/robertd2000/task-cli/internals/models"
)

func SerializeTask(task *models.Task) ([]byte, error) {
	res, err := json.Marshal(task)

	if err != nil {
		fmt.Printf("unable to marshal task: %w", err)
		return nil, err
	}

	return res, nil
}

func SaveToJSON(task []byte) error {
	err := os.WriteFile("db.json" ,task ,0755)

	if err != nil {
		fmt.Printf("unable to write file: %w", err)
		return err
	}

	return nil
}

func ReadFromJSON() ([]byte, error) {
	res, err := os.ReadFile("db.json")	

	if err != nil {
		fmt.Printf("unable to read file: %w", err)
		return nil, err
	}

	return res, nil
}

func DeserializeTask(stream []byte) (*models.Task, error) {
	task := &models.Task{}
	
	if err := json.Unmarshal(stream, task); err != nil {
		return nil, err
    }

	return task, nil
}