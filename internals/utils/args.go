package utils

import (
	"log"
	"strconv"
)

func GetDescription(args []string, idx int) string {
	if len(args) < idx+1 {
		return ""
	}

	return args[idx]
}

func GetId(args []string, idx int) int {
	if len(args) < idx+1 {
		return 0
	}

	id, err := strconv.Atoi(args[idx])
	if err != nil {
		log.Fatal("invalid id: %w", err)
	}

	return id
}