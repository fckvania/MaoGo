package commands

import (
	"mao/src/libs"
)

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:     "(setname|setnamegc|setnamegrup|setnamegroup)",
		As:       []string{"setname", "setnamegroup"},
		Tags:     "group",
		IsPrefix: true,
		IsWaitt:  true,
		IsQuerry: true,
		IsAdmin:  true,
		Exec: func(client *libs.NewClientImpl, m *libs.IMessage) {
			err := client.WA.SetGroupName(m.From, m.Querry)
			if err != nil {
				m.Reply("Mao gagal mengubah nama group")
				return
			}
			m.Reply("Berhasil mengubah nama group")
		},
	})
}
