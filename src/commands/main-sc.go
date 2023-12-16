package commands

import "mao/src/libs"

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:     "(sc|source)",
		As:       []string{"sc"},
		Tags:     "main",
		IsPrefix: true,
		Exec: func(client *libs.NewClientImpl, m *libs.IMessage) {
			m.Reply("https://github.com/fckvania/MaoGo\n\n_Free Not For Sell_")
		},
	})
}
