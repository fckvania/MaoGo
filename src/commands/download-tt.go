package commands

import (
	"mao/src/libs"
	"mao/src/libs/api"
)

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:     "(tt|tiktok)",
		As:       []string{"tiktok"},
		Tags:     "downloader",
		IsPrefix: true,
		IsQuerry: true,
		IsWaitt:  true,
		Exec: func(client *libs.NewClientImpl, m *libs.IMessage) {
			url, err := api.GetTiktokVideo(m.Querry)
			if err != nil {
				m.Reply(err.Error())
				return
			}

			bytes, err := client.GetBytes(url)
			if err != nil {
				m.Reply(err.Error())
				return
			}
			client.SendVideo(m.From, bytes, "", m.ID)
		},
	})
}
