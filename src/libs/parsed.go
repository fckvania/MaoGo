package libs

import (
	"os"
	"strings"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
)

func NewSmsg(mess *events.Message, sock *NewClientImpl) *IMessage {
	var command string
	var media whatsmeow.DownloadableMessage
	var isOwner = false
	var owner []string
	botNum, _ := sock.ParseJID(sock.WA.Store.ID.User)
	quotedMsg := mess.Message.GetExtendedTextMessage().GetContextInfo().GetQuotedMessage()
	owner = append(owner, os.Getenv("Owner_Number"))
	owner = append(owner, botNum.String())

	for _, own := range owner {
		if own == mess.Info.Sender.String() {
			isOwner = true
		}
	}

	if pe := mess.Message.GetExtendedTextMessage().GetText(); pe != "" {
		command = pe
	} else if pe := mess.Message.GetConversation(); pe != "" {
		command = pe
	} else if pe := mess.Message.GetImageMessage().GetCaption(); pe != "" {
		command = pe
	} else if pe := mess.Message.GetVideoMessage().GetCaption(); pe != "" {
		command = pe
	}

	if quotedMsg != nil && (quotedMsg.ImageMessage != nil || quotedMsg.VideoMessage != nil) {
		if msg := quotedMsg.GetImageMessage(); msg != nil {
			media = msg
		} else if msg := quotedMsg.GetVideoMessage(); msg != nil {
			media = msg
		}
	} else if mess.Message != nil && (mess.Message.ImageMessage != nil || mess.Message.VideoMessage != nil) {
		if msg := mess.Message.GetImageMessage(); msg != nil {
			media = msg
		} else if msg := mess.Message.GetVideoMessage(); msg != nil {
			media = msg
		}
	} else {
		media = nil
	}

	return &IMessage{
		From:     mess.Info.Chat,
		PushName: mess.Info.PushName,
		IsOwner:  isOwner,
		Querry:   strings.Join(strings.Split(command, " ")[1:], ` `),
		Command:  strings.Split(command, " ")[0],
		Media:    media,
		IsImage: func() bool {
			if mess.Message.GetImageMessage() != nil {
				return true
			} else {
				return false
			}
		}(),
		ContextInfo: &waProto.ContextInfo{
			StanzaId:      &mess.Info.ID,
			Participant:   proto.String(mess.Info.Sender.String()),
			QuotedMessage: mess.Message,
		},
		IsQuotedImage: func() bool {
			if quotedMsg.GetImageMessage() != nil {
				return true
			} else {
				return false
			}
		}(),
		Reply: func(text string) {
			sock.SendText(mess.Info.Chat, text, &waProto.ContextInfo{
				StanzaId:      &mess.Info.ID,
				Participant:   proto.String(mess.Info.Sender.String()),
				QuotedMessage: mess.Message,
			})
		},
	}
}
