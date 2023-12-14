package libs

import (
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
)

type NewClientImpl struct {
	WA *whatsmeow.Client
}

type ICommand struct {
	Name        string
	As          []string
	Description string
	Tags        string
	IsPrefix    bool
	IsOwner     bool
	IsMedia     bool
	IsQuerry    bool
	IsGroup     bool
	IsAdmin     bool
	After       func(client *NewClientImpl, m *IMessage)
	Exec        func(client *NewClientImpl, m *IMessage)
}

type IMessage struct {
	From          types.JID
	IsBot         bool
	Sender        types.JID
	PushName      string
	IsOwner       bool
	IsAdmin       bool
	IsGroup       bool
	Querry        string
	Command       string
	IsImage       bool
	IsVideo       bool
	IsQuotedImage bool
	IsQuotedVideo bool
	Media         whatsmeow.DownloadableMessage
	ContextInfo   *waProto.ContextInfo
	Reply         func(text string)
}
