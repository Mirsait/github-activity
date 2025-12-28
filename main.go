package main

import (
	"fmt"
	"os"

	"github.com/Mirsait/github-activity/models"
	"github.com/Mirsait/github-activity/network"
	"github.com/Mirsait/github-activity/storage"
)

type Activity = models.Activity
type GithubEvent = models.GithubEvent

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
			fmt.Println(err.Error())
			return
		}
		activities = append(activities, result...)
	} else {
		result, err := network.GetGithubActivities(username)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		activities = append(activities, result...)
		err = storage.Save(filename, activities)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	history := buildMap(activities)
	for repo, events := range history {
		for ghEvent, count := range events {
			text := ghEvent.GetText(repo, count)
			fmt.Println(text)
		}
	}
}

// repo - event - count
type History = map[string]map[GithubEvent]uint

func buildMap(activities []Activity) History {
	history := make(History)
	for _, act := range activities {
		repo := act.Repo.Name
		event, _ := act.GetGithubEvent()
		if history[repo] == nil {
			history[repo] = make(map[GithubEvent]uint)
		}
		c := history[repo][event]
		history[repo][event] = c + 1
	}
	return history
}
