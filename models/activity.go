package models

import (
	"encoding/json"
)

type Event string

type Activity struct {
	EventType string          `json:"type"`
	Repo      Repository      `json:"repo"`
	Actor     Actor           `json:"actor"`
	CreatedAt string          `json:"created_at"`
	Payload   json.RawMessage `json:"payload"`
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
