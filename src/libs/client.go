package libs

import (
	"context"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

type NewClientImpl struct {
	WA *whatsmeow.Client
}

func NewClient(client *whatsmeow.Client) *NewClientImpl {
	return &NewClientImpl{
		WA: client,
	}
}

func (client *NewClientImpl) SendText(from types.JID, txt string, opts *waProto.ContextInfo) {
	client.WA.SendMessage(context.Background(), from, &waProto.Message{
		ExtendedTextMessage: &waProto.ExtendedTextMessage{
			Text:        proto.String(txt),
			ContextInfo: opts,
		},
	})
}

func (client *NewClientImpl) SendWithNewsLestter(from types.JID, text string, newjid string, newserver int32, name string, opts *waProto.ContextInfo) {
	client.SendText(from, text, &waProto.ContextInfo{
		ForwardedNewsletterMessageInfo: &waProto.ForwardedNewsletterMessageInfo{
			NewsletterJid:   proto.String(newjid),
			NewsletterName:  proto.String(name),
			ServerMessageId: proto.Int32(newserver),
		},
		IsForwarded:   proto.Bool(true),
		StanzaId:      opts.StanzaId,
		Participant:   opts.Participant,
		QuotedMessage: opts.QuotedMessage,
	})
}
