package commands

import (
	"fmt"
	"mao/src/libs"
)

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:     "(setpicture|setpp|setppgrup|setppgroup)",
		As:       []string{"setpp", "setppgroup"},
		Tags:     "group",
		IsPrefix: true,
		IsWaitt:  true,
		IsAdmin:  true,
		Exec: func(client *libs.NewClientImpl, m *libs.IMessage) {
			if !(m.IsImage || m.IsQuotedImage) {
				m.Reply("Balas/kirim pesan image!")
				return
			}
			var img []byte
			var err error
			if m.IsQuotedImage {
				img, err = client.WA.Download(m.QuotedMsg.QuotedMessage.GetImageMessage())
			} else {
				img, err = client.WA.Download(m.Media)
			}
			if err != nil {
				m.Reply("terjadi kesalahan saat mendownload image")
				fmt.Println(err.Error())
				return
			}
			if img == nil {
				m.Reply("Balas/kirim pesan image")
				return
			}
			_, err = client.WA.SetGroupPhoto(m.From, img)
			if err != nil {
				m.Reply("Gagal mengubah foto group")
				fmt.Println(err.Error())
			}
		},
	})
}
