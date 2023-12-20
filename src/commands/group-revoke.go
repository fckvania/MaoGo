package commands

import (
	"fmt"
	"mao/src/libs"
)

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:     "(revoke|resetlink)",
		As:       []string{"revoke", "resetlink"},
		Tags:     "group",
		IsPrefix: true,
		IsWaitt:  true,
		IsAdmin:  true,
		Exec: func(client *libs.NewClientImpl, m *libs.IMessage) {
			resp, err := client.WA.GetGroupInviteLink(m.From, true)
			if err != nil {
				m.Reply("Mao tidak bisa mereset group link :(")
			} else {
				m.Reply(fmt.Sprintf("Link group baru: %s", resp))
			}
		},
	})
}
