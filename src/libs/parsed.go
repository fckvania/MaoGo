package libs

import (
	"os"
	"strings"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
)

func NewSmsg(mess *events.Message, sock *NewClientImpl) *IMessage {
	var command string
	var isOwner = false
	var owner []string
	botNum, _ := sock.ParseJID(sock.WA.Store.ID.User)
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

	return &IMessage{
		From:     mess.Info.Chat,
		PushName: mess.Info.PushName,
		IsOwner:  isOwner,
		Querry:   strings.Join(strings.Split(command, " ")[1:], ` `),
		Command:  strings.Split(command, " ")[0],
		ContextInfo: &waProto.ContextInfo{
			StanzaId:      &mess.Info.ID,
			Participant:   proto.String(mess.Info.Sender.String()),
			QuotedMessage: mess.Message,
		},
		Reply: func(text string) {
			sock.SendText(mess.Info.Chat, text, &waProto.ContextInfo{
				StanzaId:      &mess.Info.ID,
				Participant:   proto.String(mess.Info.Sender.String()),
				QuotedMessage: mess.Message,
			})
		},
	}
}
