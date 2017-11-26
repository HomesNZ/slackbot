package main

import (
	git "./api/git"
	slack "./api/slack"
)

// "github.com/jimsrush/slackbot/handlers"

func main() {
	git.InitGit()
	slack.InitSlack()
}
