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
	Exec        func(client *NewClientImpl, m *IMessage)
}

type IMessage struct {
	From        types.JID
	PushName    string
	IsOwner     bool
	Querry      string
	Reply       func(text string)
	Command     string
	ContextInfo *waProto.ContextInfo
}
