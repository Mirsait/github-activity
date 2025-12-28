package models

import (
	"encoding/json"
	"fmt"
	"unicode"
)

// https://docs.github.com/en/rest/using-the-rest-api/github-event-types?apiVersion=2022-11-28

// "DeleteEvent", payload -> ref_type = branch | tag
// "CreateEvent", payload -> ref_type = branch | tag | repository
//
// "IssuesEvent", payload -> action = opened | closed | reopened
// "PullRequestEvent", payload -> action = opened | closed | merged | reopened | assigned | unassigned | labeled | unlabeled
// "PullRequestReviewEvent", payload -> action = created | updated | dismissed
//
// "DiscussionEvent"
// "IssueCommentEvent"
// "ForkEvent"
// "GollumEvent"
// "MemberEvent"
// "PublicEvent"
// "PullRequestReviewCommentEvent"
// "PushEvent"
// "ReleaseEvent"
// "CommitCommentEvent"
// "WatchEvent"

type GithubEvent struct {
	Type    string
	Payload string
}

func (a *Activity) GetGithubEvent() (GithubEvent, error) {
	event := GithubEvent{Type: a.EventType}
	switch a.EventType {
	case "IssuesEvent", "PullRequestEvent", "PullRequestReviewEvent":
		var payload PayloadAction
		err := json.Unmarshal(a.Payload, &payload)
		event.Payload = payload.Action
		return event, err
	case "DeleteEvent", "CreateEvent":
		var payload PayloadRef
		err := json.Unmarshal(a.Payload, &payload)
		event.Payload = payload.RefType
		return event, err
	case "DiscussionEvent", "IssueCommentEvent", "ForkEvent", "GollumEvent",
		"MemberEvent", "PublicEvent", "PullRequestReviewCommentEvent",
		"PushEvent", "ReleaseEvent", "CommitCommentEvent", "WatchEvent":
		return event, nil
	default:
		return GithubEvent{}, fmt.Errorf("unknown event type: %s", a.EventType)
	}
}

func (e *GithubEvent) GetText(repo string, count uint) string {
	typeEvent := e.Type
	payload := e.Payload
	switch typeEvent {
	case "DeleteEvent":
		return deleteMsg(repo, payload, count)
	case "CreateEvent":
		return createMsg(repo, payload, count)
	case "IssuesEvent":
		return issueMsg(repo, payload, count)
	case "PullRequestEvent":
		return pullRequestMsg(repo, payload, count)
	case "PullRequestReviewEvent":
		return pullRequestReviewMsg(repo, payload, count)
	case "DiscussionEvent":
		return createDiscussion(repo, count)
	case "IssueCommentEvent":
		return createIssueComment(repo, count)
	case "ForkEvent":
		return forkRepository(repo, count)
	case "GollumEvent":
		return createdWiki(repo, count)
	case "MemberEvent":
		return fmt.Sprintf("- Added to repository collaborators in repository %s", repo)
	case "PublicEvent":
		return fmt.Sprintf("- Private repository %s is made public", repo)
	case "PullRequestReviewCommentEvent":
		return fmt.Sprintf("- Created %d comments for pull request in repository %s", count, repo)
	case "PushEvent":
		return pushMsg(repo, count)
	case "ReleaseEvent":
		return releaseMsg(repo, count)
	case "CommitCommentEvent":
		return fmt.Sprintf("- Created %d commit comments in repository %s", count, repo)
	case "WatchEvent":
		return fmt.Sprintf("- Stared a repository %s", repo)
	}
	return fmt.Sprintf("Undefinded event: %s", typeEvent)
}

func createdWiki(repo string, count uint) string {
	if count == 1 {
		return fmt.Sprintf("- In repository %s is created or updated a wiki page", repo)
	}
	return fmt.Sprintf("- In repository %s is created or updated %d wiki pages", repo, count)
}

func forkRepository(repo string, count uint) string {
	if count == 1 {
		return fmt.Sprintf("- Forked one repository %s", repo)
	}
	return fmt.Sprintf("- Forked %d repositories", count)
}

func createIssueComment(repo string, count uint) string {
	if count == 1 {
		return fmt.Sprintf("- Created an issue or pull request comment in repository %s", repo)
	}
	return fmt.Sprintf("- Created %d issue or pull request comments in repository %s", count, repo)
}

func createDiscussion(repo string, count uint) string {
	if count == 1 {
		return fmt.Sprintf("- A discussion is created in a repository %s", repo)
	}
	return fmt.Sprintf("- %d discussions is created in a repository %s", count, repo)
}

func deleteMsg(repo, payload string, count uint) string {
	if count == 1 {
		return fmt.Sprintf("- Deleted one %s in %s", payload, repo)
	}
	return fmt.Sprintf("- Deleted %d %s(e)s in %s", count, payload, repo)
}

func createMsg(repo, payload string, count uint) string {
	if count == 1 {
		return fmt.Sprintf("- Created a new %s in %s", payload, repo)
	}
	return fmt.Sprintf("- Created %d %s(e)s in %s", count, payload, repo)
}

func issueMsg(repo, action string, count uint) string {
	if count == 1 {
		return fmt.Sprintf("- %s a new issue in %s", action, repo)
	}
	return fmt.Sprintf("- %s %d issues in %s", action, count, repo)
}

func capitalize(s string) string {
	if len(s) == 0 {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func pullRequestMsg(repo, action string, count uint) string {
	title := capitalize(action)
	if count == 1 {
		return fmt.Sprintf("- %s a pull request in %s", title, repo)
	}
	return fmt.Sprintf("- %s %d pull requests in %s", title, count, repo)
}

func pullRequestReviewMsg(repo, action string, count uint) string {
	title := capitalize(action)
	if count == 1 {
		return fmt.Sprintf("- %s a pull request review in %s", title, repo)
	}
	return fmt.Sprintf("- %s %d pull request reviews in %s", title, count, repo)
}

func pushMsg(repo string, count uint) string {
	if count == 1 {
		return fmt.Sprintf("- Pushed a new commit in %s", repo)
	}
	return fmt.Sprintf("- Pushed  %d commits in %s", count, repo)
}

func releaseMsg(repo string, count uint) string {
	if count == 1 {
		return fmt.Sprintf("- Made a release in %s", repo)
	}
	return fmt.Sprintf("- Made %d releases in %s", count, repo)
}
