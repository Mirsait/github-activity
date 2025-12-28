package network

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Mirsait/github-activity/models"
)

type Activity = models.Activity

func GetGithubActivities(name string) ([]Activity, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// only last 100 events
	url := fmt.Sprintf("http://api.github.com/users/%s/events?per_page=100", name)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	request.Header.Add("Accept", "application/vnd.github+json")

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	var result []Activity
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("parse JSON: %w", err)
	}
	return result, nil
}
