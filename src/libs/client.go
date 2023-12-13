package libs

import (
	"context"
	"strings"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

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

func (simp *NewClientImpl) ParseJID(arg string) (types.JID, bool) {
	if arg[0] == '+' {
		arg = arg[1:]
	}
	if !strings.ContainsRune(arg, '@') {
		return types.NewJID(arg, types.DefaultUserServer), true
	} else {
		recipient, err := types.ParseJID(arg)
		if err != nil {
			return recipient, false
		} else if recipient.User == "" {
			return recipient, false
		}
		return recipient, true
	}
}
