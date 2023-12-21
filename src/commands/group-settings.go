package commands

import (
	"fmt"
	"mao/src/libs"
	"strings"
)

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:     "(opengc|opengroup|closegc|closegroup)",
		As:       []string{"opengc", "closegc"},
		Tags:     "group",
		IsPrefix: true,
		IsWaitt:  true,
		IsAdmin:  true,
		IsGroup:  true,
		Exec: func(client *libs.NewClientImpl, m *libs.IMessage) {
			err := client.WA.SetGroupLocked(m.From, strings.Contains(m.Command, "close"))
			if err != nil {
				m.Reply("Error")
				fmt.Println(err.Error())
			}
		},
	})
}
