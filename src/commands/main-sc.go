package commands

import "mao/src/libs"

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:     "sc",
		Tags:     "main",
		IsPrefix: true,
		Exec: func(client *libs.NewClientImpl, m *libs.IMessage) {
			m.Reply("https://github.com/fckvania/MaoGo")
		},
	})
}
