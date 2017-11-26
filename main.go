package main

import (
	"fmt"
	"os"
	"strings"

	handlers "./handlers"
	// "github.com/jimsrush/slackbot/handlers"
	"github.com/nlopes/slack"
)

func main() {

	token := os.Getenv("SLACK_TOKEN")
	api := slack.New(token)
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Print("Event Received: ")
			switch ev := msg.Data.(type) {

			case *slack.MessageEvent:
				fmt.Printf("Message: %v\n", ev)
				info := rtm.GetInfo()
				prefix := fmt.Sprintf("<@%s> ", info.User.ID)

				if ev.User != info.User.ID && strings.HasPrefix(ev.Text, prefix) {
					response := handlers.HandleMessage(ev.Text)
					rtm.SendMessage(rtm.NewOutgoingMessage(response, ev.Channel))

				}

			default:
				//Take no action
			}
		}
	}
}
