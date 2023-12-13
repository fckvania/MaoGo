package libs

import "go.mau.fi/whatsmeow"

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
