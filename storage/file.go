package storage

import (
	"belajar-golang/models"
	"encoding/json"
	"errors"
	"io/fs"
	"os"
)

var DataFile = "tasks.json"

// LoadTasks membaca semua task dari file JSON
func LoadTasks() ([]models.Task, error) {
	var tasks []models.Task

	data, err := os.ReadFile(DataFile)
	if errors.Is(err, fs.ErrNotExist) {
		return tasks, nil // file belum ada
	}
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return tasks, nil
	}
	return tasks, json.Unmarshal(data, &tasks)
}

// SaveTasks menyimpan semua task ke file JSON
func SaveTasks(tasks []models.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(DataFile, data, 0o644)
}
