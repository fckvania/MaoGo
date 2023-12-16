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
				return
			}

			client.SendVideo(m.From, client.GetBytes(url), "", nil)
		},
	})
}
