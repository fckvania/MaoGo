package commands

import (
	"mao/src/libs"
	"strconv"
)

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:     "(setmode|mode)",
		As:       []string{"setmode"},
		Tags:     "owner",
		IsPrefix: true,
		IsOwner:  true,
		Exec: func(client *libs.NewClientImpl, m *libs.IMessage) {
			if m.Querry == "public" {
				client.Public = true
				m.Reply("Public Mode: " + strconv.FormatBool(client.Public))
			} else {
				client.Public = false
				m.Reply("Public Mode: " + strconv.FormatBool(client.Public))
			}
		},
	})
}
