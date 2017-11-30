package handlers

import (
	"fmt"
	"strings"

	git "github.com/jimsrush/slackbot/api/git"
)

//HandleMessage is the handler for a Real Time Slack message Struct
func HandleMessage(s string) string {
	fmt.Print("message is", strings.Split(s, " "))
	message := strings.Split(s, " ")

	if message[1] == "pull" {
		if len(message) == 2 {
			return git.GetPullRequestData()
		}
		if len(message) > 2 {
			return git.GetPullRequestDataByRepo(message[2])
		}

		return "Should I get some requests"
	}

	// go git.GetPullRequestData()
	return "no"
	// return "Hey"
}
