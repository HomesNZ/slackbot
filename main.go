package main

import (
	git "github.com/jimsrush/slackbot/api/git"
	slack "github.com/jimsrush/slackbot/api/slack"
)

func main() {
	git.InitGit()
	slack.InitSlack()
}
