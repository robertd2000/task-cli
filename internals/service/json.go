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
	}

	return nil
}