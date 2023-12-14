package commands

import (
	"fmt"
	"mao/src/libs"
)

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:     "(channelinfo|ci)",
		As:       []string{"ci"},
		Tags:     "main",
		IsPrefix: true,
		IsQuerry: true,
		Exec: func(client *libs.NewClientImpl, m *libs.IMessage) {
			key, err := client.WA.GetNewsletterInfoWithInvite(m.Querry)
			if err != nil {
				m.Reply("Mao Tidak Tau Ya.")
				return
			}

			m.Reply(fmt.Sprintf("Channel Information\nLink: %s\nID: %s\nName: %v\nFollowers: %v\n\nDescription: %v\nCreate At: %v", m.Querry, key.ID, key.ThreadMeta.Name.Text, key.ThreadMeta.SubscriberCount, key.ThreadMeta.Description.Text, key.ThreadMeta.CreationTime))
		},
	})
}
