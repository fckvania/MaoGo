package commands

import (
	"fmt"
	"mao/src/libs"
	"os/exec"
)

func _exec(client *libs.NewClientImpl, m *libs.IMessage) {
	out, err := exec.Command("bash", "-c", m.Querry).Output()
	if err != nil {
		m.Reply(fmt.Sprintf("%v", err))
		return
	}
	m.Reply(string(out))

}

func init() {
	handler := libs.ICommand{
		Name:     "$",
		Tags:     "owner",
		IsPrefix: false,
		IsOwner:  true,
		Exec:     _exec,
	}
	libs.NewCommands().Add(&handler)
}
