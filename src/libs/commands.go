package libs

import (
	"regexp"
	"strings"
)

var lists []ICommand

func NewCommands(cmd *ICommand) {
	lists = append(lists, *cmd)
}

func GetList() []ICommand {
	return lists
}

func Get(c *NewClientImpl, m *IMessage) {
	var prefix string
	pattern := regexp.MustCompile(`[?!.#]`)
	for _, f := range pattern.FindAllString(m.Command, -1) {
		prefix = f
	}
	for _, cmd := range lists {
		if cmd.After != nil {
			cmd.After(c, m)
		}
		re := regexp.MustCompile(`^` + cmd.Name + `$`)
		if valid := len(re.FindAllString(strings.ReplaceAll(m.Command, prefix, ""), -1)) > 0; valid {
			var cmdWithPref bool
			var cmdWithoutPref bool
			if cmd.IsPrefix && (prefix != "" && strings.HasPrefix(m.Command, prefix)) {
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

			//Read Command
			//c.WA.MarkRead([]string{m.StanzaId}, time.Now(), m.From, m.Sender)

			//Checking
			if cmd.IsOwner && !m.IsOwner {
				continue
			}

			if cmd.IsMedia && m.Media == nil {
				m.Reply("Media Di Butuhkan")
				continue
			}

			if cmd.IsQuerry && m.Querry == "" {
				m.Reply("Querry Di Butuhkan")
				continue
			}

			if cmd.IsGroup && !m.IsGroup {
				m.Reply("Hanya Khusus Group")
				continue
			}

			if cmd.IsPrivate && m.IsGroup {
				m.Reply("Hanya Khusus Private")
				continue
			}

			if (m.IsGroup && cmd.IsAdmin) && !m.IsAdmin {
				m.Reply("Hanya Khusus Admin")
				continue
			}

			if (m.IsGroup && cmd.IsBotAdmin) && !m.IsBotAdmin {
				m.Reply("Jadikan Bot Sebagai Admin")
				continue
			}

			if cmd.IsWaitt {
				m.Reply("Tunggu Sebentar Ya")
			}

			cmd.Exec(c, m)
		}
	}
}
