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
	Description string
	Tags        string
	IsPrefix    bool
	IsOwner     bool
	IsMedia     bool
	IsQuerry    bool
	Exec        func(client *NewClientImpl, m *IMessage)
}

type IMessage struct {
	From          types.JID
	PushName      string
	IsOwner       bool
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
