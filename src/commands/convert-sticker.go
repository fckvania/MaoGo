package commands

import (
	"context"
	"fmt"
	"mao/src/libs"
	"mao/src/libs/api"
	"net/http"

	"go.mau.fi/whatsmeow"
	"google.golang.org/protobuf/proto"

	waProto "go.mau.fi/whatsmeow/binary/proto"
)

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:     "sticker",
		Tags:     "convert",
		IsPrefix: true,
		IsMedia:  true,
		Exec: func(client *libs.NewClientImpl, m *libs.IMessage) {
			data, _ := client.WA.Download(m.Media)

			s := api.StickerApi(&api.Sticker{
				File: data,
				Tipe: func() api.MediaType {
					if m.IsImage || m.IsQuotedImage {
						return api.IMAGE
					} else {
						return api.VIDEO
					}
				}(),
			}, &api.MetadataSticker{
				Author:    "github.com/fckvania",
				Pack:      "Mao",
				KeepScale: true,
				Removebg:  "false",
				Circle: func() bool {
					if m.Querry == "-c" {
						return true
					} else {
						return false
					}
				}(),
			})

			stc := s.Build()

			uploaded, err := client.WA.Upload(context.Background(), stc, whatsmeow.MediaImage)
			if err != nil {
				fmt.Printf("Failed to upload file: %v\n", err)
			}

			client.WA.SendMessage(context.Background(), m.From, &waProto.Message{
				StickerMessage: &waProto.StickerMessage{
					Url:           proto.String(uploaded.URL),
					DirectPath:    proto.String(uploaded.DirectPath),
					MediaKey:      uploaded.MediaKey,
					Mimetype:      proto.String(http.DetectContentType(stc)),
					FileEncSha256: uploaded.FileEncSHA256,
					FileSha256:    uploaded.FileSHA256,
					FileLength:    proto.Uint64(uint64(len(data))),
					ContextInfo:   m.ContextInfo,
				},
			})

		},
	})
}
