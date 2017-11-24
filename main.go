package main

import (
	"os"

	"github.com/nlopes/slack"
)

func main() {

	token := os.Getenv("SLACK_TOKEN")
	api := slack.New(token)
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for {
		msg := <-rtm.IncomingEvents
		// fmt.Println("Message is: ", msg)
		handlers.HandleMessage(msg)
	}
}
