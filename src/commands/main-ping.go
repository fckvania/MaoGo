package commands

import "mao/src/libs"

func init() {
	handler := libs.ICommand{
		Name:     "ping",
		Tags:     "main",
		IsPrefix: true,
		Exec: func(client *libs.NewClientImpl, m *libs.IMessage) {
			m.Reply("Pong!")
		},
	}
	libs.NewCommands().Add(&handler)
}
