package commands

import (
	"fmt"
	"mao/src/libs"
	"strings"
)

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:     "(opengc|opengroup|closegc|closegroup)",
		As:       []string{"opengc", "opengroup", "closegc", "closegroup"},
		Tags:     "group",
		IsPrefix: true,
		IsWaitt:  true,
		IsAdmin:  true,
		Exec: func(client *libs.NewClientImpl, m *libs.IMessage) {
			err := client.WA.SetGroupAnnounce(m.From, strings.Contains(m.Command, "open"))
			if err != nil {
				m.Reply("Error")
				fmt.Println(err.Error())
			}
		},
	})
}
