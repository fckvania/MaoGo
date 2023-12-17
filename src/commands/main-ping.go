package commands

import (
	"mao/src/libs"
	"time"
)

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:     "ping",
		As:       []string{"ping"},
		Tags:     "main",
		IsPrefix: true,
		Exec: func(client *libs.NewClientImpl, m *libs.IMessage) {
			start := time.Now()
			m.Reply("Pong!")
			m.Reply("Speed: " + time.Since(start).String())
		},
	})
}
