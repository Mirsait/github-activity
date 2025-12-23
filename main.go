package main

import (
	"fmt"
	"os"

	"github.com/Mirsait/github-activity/models"
	"github.com/Mirsait/github-activity/network"
	"github.com/Mirsait/github-activity/storage"
)

type Activity = models.Activity

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage `github-activity Username`")
		return
	}
	username := args[0]
	filename := fmt.Sprintf("data/%s.json", username)

	var activities []Activity
	if storage.Exists(filename) {
		result, err := storage.Load(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		activities = append(activities, result...)
	} else {
		result, err := network.GetGithubActivities(username)
		if err != nil {
			fmt.Println(err)
		}
		activities = append(activities, result...)
		err = storage.Save(filename, activities)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	history := buildMap(activities)
	for outKey, innerMap := range history {
		for inKey, value := range innerMap {
			fmt.Printf("- %d %s in %s\n", value, inKey, outKey)
		}
	}
}

type History = map[string]map[string]uint

func buildMap(activities []Activity) History {
	history := make(History)
	for _, act := range activities {
		repo := act.Repo.Name
		eventType := act.EventType
		if history[repo] == nil {
			history[repo] = make(map[string]uint)
		}
		c := history[repo][eventType]
		history[repo][eventType] = c + 1
	}
	return history
}
