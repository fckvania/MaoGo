package libs

import (
	"strings"
)

type ICommand struct {
	Name        string
	Description string
	Tags        string
	IsPrefix    bool
	IsOwner     bool
	Exec        func(client *NewClientImpl, m *IMessage)
}

var lists []ICommand

type newCmd struct {
	cmd ICommand
}

func NewCommands() *newCmd {
	return &newCmd{}
}

func (b *newCmd) Add(cmd *ICommand) {
	lists = append(lists, *cmd)
}

func (b *newCmd) GetList() []ICommand {
	return lists
}

func (b *newCmd) Get(c *NewClientImpl, m *IMessage) {
	prefix := "#"
	for _, cmd := range lists {
		if cmd.Name == strings.ReplaceAll(m.Command, prefix, "") {
			var cmdWithPref bool
			var cmdWithoutPref bool

			if cmd.IsPrefix && strings.HasPrefix(m.Command, prefix) {
				cmdWithPref = true
			} else {
				cmdWithPref = false
			}

			if !cmd.IsPrefix {
				cmdWithoutPref = true
			} else {
				cmdWithoutPref = false
			}

			if !cmdWithPref && !cmdWithoutPref {
				continue
			}

			if cmd.IsOwner && !m.IsOwner {
				continue
			}

			cmd.Exec(c, m)
		}
	}
}
