package models

import (
	"encoding/json"
	"fmt"
)

type Event string

type Activity struct {
	EventType string          `json:"type"`
	Repo      Repository      `json:"repo"`
	Actor     Actor           `json:"actor"`
	CreatedAt string          `json:"created_at"`
	Payload   json.RawMessage `json:"payload"`
}

func (a *Activity) GetPayload() (any, error) {
	switch a.EventType {
	case "IssuesEvent", "PullRequestEvent", "PullRequestReviewEvent":
		var payload PayloadAction
		err := json.Unmarshal(a.Payload, &payload)
		return payload, err
	case "DeleteEvent", "CreateEvent":
		var payload PayloadRef
		err := json.Unmarshal(a.Payload, &payload)
		return payload, err
	default:
		return nil, fmt.Errorf("unknown event type: %s", a.EventType)
	}
}

type Actor struct {
	Login string `json:"login"`
	Url   string `json:"url"`
}

type Repository struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PayloadAction struct {
	Action string `json:"action"`
}
type PayloadRef struct {
	RefType string `json:"ref_type"`
}
