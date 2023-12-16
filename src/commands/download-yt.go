package commands

import (
	"fmt"
	"mao/src/libs"
	"mao/src/libs/api"
	"regexp"
)

type method string

const (
	audio string = "Audio"
	video string = "Video"
)

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:     "(ytmp4|ytmp3)",
		As:       []string{"ytmp4", "ytmp3"},
		Tags:     "downloader",
		IsPrefix: true,
		IsQuerry: true,
		IsWaitt:  true,
		Exec: func(client *libs.NewClientImpl, m *libs.IMessage) {
			yt, err := api.YoutubeDL(m.Querry)
			if err != nil {
				m.Reply(err.Error())
				return
			}

			caption := fmt.Sprintf("*Title*: %s\n*Author*: %s", yt.Info.Title, yt.Info.Author)

			if reg, _ := regexp.MatchString(`(ytmp3)`, m.Command); reg {
				build, err := yt.Link.Audio[0].Url()
				if err != nil {
					m.Reply(err.Error())
					return
				}
				client.SendDocument(m.From, client.GetBytes(build), fmt.Sprintf("%s.mp3", yt.Info.Title), caption, m.ContextInfo)
			} else {
				build, err := yt.Link.Video[0].Url()
				if err != nil {
					m.Reply(err.Error())
					return
				}
				client.SendVideo(m.From, client.GetBytes(build), caption, m.ContextInfo)
			}
		},
	})
}
