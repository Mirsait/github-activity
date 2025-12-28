package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Mirsait/github-activity/models"
)

func Save(filename string, data []models.Activity) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("encode: %w", err)
	}
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("write data: %w", err)
	}
	return nil
}
