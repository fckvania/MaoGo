package commands

import (
	"mao/src/libs"
)

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:     "(setdesc|setdeskripsi|setdesk)",
		As:       []string{"setdesc", "setdeskripsi"},
		Tags:     "group",
		IsPrefix: true,
		IsWaitt:  true,
		IsQuerry: true,
		IsAdmin:  true,
		Exec: func(client *libs.NewClientImpl, m *libs.IMessage) {
			err := client.WA.SetGroupTopic(m.From, "", "", m.Querry)
			if err != nil {
				m.Reply("Mao gagal mengubah deskripsi group")
				return
			}
			m.Reply("Berhasil mengubah deskripsi group")
		},
	})
}
