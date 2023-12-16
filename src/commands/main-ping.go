package commands

import (
	"mao/src/libs"
)

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:     "ping",
		As:       []string{"ping"},
		Tags:     "main",
		IsPrefix: true,
		IsAdmin:  true,
		Exec: func(client *libs.NewClientImpl, m *libs.IMessage) {
			m.Reply("Pong!")
		},
	})
}
