package commands

import (
	"fmt"
	"mao/src/libs"
	"os/exec"
)

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:     "$",
		Tags:     "owner",
		IsPrefix: false,
		IsOwner:  true,
		Exec: func(client *libs.NewClientImpl, m *libs.IMessage) {
			out, err := exec.Command("bash", "-c", m.Querry).Output()
			if err != nil {
				m.Reply(fmt.Sprintf("%v", err))
				return
			}
			m.Reply(string(out))
		},
	})
}
