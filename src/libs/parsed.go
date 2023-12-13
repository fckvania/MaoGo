package libs

import (
	"os"
	"strings"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
)

type IMessage struct {
	From        types.JID
	PushName    string
	IsOwner     bool
	Querry      string
	Reply       func(text string)
	Command     string
	ContextInfo *waProto.ContextInfo
}

func NewSmsg(mess *events.Message, sock *NewClientImpl) *IMessage {
	var command string
	var isOwner = false
	if mess.Info.Sender.String() == os.Getenv("Owner_Number") {
		isOwner = true
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
