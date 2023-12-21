package commands

import (
	"fmt"
	"mao/src/libs"
)

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:       "(linkgroup|linkgrup|linkgc)",
		As:         []string{"linkgroup"},
		Tags:       "group",
		IsPrefix:   true,
		IsWaitt:    true,
		IsGroup:    true,
		IsBotAdmin: true,
		Exec: func(client *libs.NewClientImpl, m *libs.IMessage) {
			resp, err := client.WA.GetGroupInviteLink(m.From, false)
			if err != nil {
				m.Reply("Gagal mendapatkan link group.")
			} else {
				m.Reply(fmt.Sprintf("Link group: %s", resp))
			}
		},
	})
}
