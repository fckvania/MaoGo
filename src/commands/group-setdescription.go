package commands

import (
	"mao/src/libs"
)

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:       "(setdesc|setdeskripsi|setdesk)",
		As:         []string{"setdesc"},
		Tags:       "group",
		IsPrefix:   true,
		IsWaitt:    true,
		IsQuerry:   true,
		IsAdmin:    true,
		IsGroup:    true,
		IsBotAdmin: true,
		Exec: func(client *libs.NewClientImpl, m *libs.IMessage) {
			err := client.WA.SetGroupTopic(m.From, "", "", m.Querry)
			if err != nil {
				m.Reply("Gagal mengubah deskripsi group")
				return
			}
			m.Reply("Berhasil mengubah deskripsi group")
		},
	})
}
