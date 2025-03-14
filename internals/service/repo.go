package service

import (
	"encoding/json"
	"os"

	"github.com/robertd2000/task-cli/internals/models"
)

func SerializeToJSON[T any](data T) ([]byte, error) {
	res, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func SaveToJSON(task []byte) error {
	err := os.WriteFile("db.json" ,task ,0755)

	if err != nil {
		return err
	}

	return nil
}

func ReadFromJSON() ([]byte, error) {
	res, err := os.ReadFile("db.json")	

	if err != nil {
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