package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Mirsait/github-activity/models"
)

func Load(filename string) ([]models.Activity, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.WriteFile(filename, []byte("[]"), 0644); err != nil {
				return nil, fmt.Errorf("create file: %w", err)
			}
		}
		return nil, fmt.Errorf("read file: %w", err)
	}

	var result []models.Activity
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("parse JSON: %w", err)
	}
	return result, nil
}
