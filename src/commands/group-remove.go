package commands

import (
	"fmt"
	"mao/src/libs"
	"strings"

	"github.com/nyaruka/phonenumbers"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
)

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:     "(kick)",
		As:       []string{"kick"},
		Tags:     "group",
		IsPrefix: true,
		IsGroup:  true,
		IsAdmin:  true,
		Exec: func(client *libs.NewClientImpl, m *libs.IMessage) {
			var ujid []types.JID
			var ok error
			// kalau error bilang :v
			// fitur belum dicoba :v
			if m.QuotedMsg != nil {
				if m.QuotedMsg.MentionedJid != nil {
					ajid := m.QuotedMsg.MentionedJid
					ujid = make([]types.JID, len(ajid))
					for i, a := range ajid {
						ujid[i], ok = types.ParseJID(a)
						if ok != nil {
							return
						}
					}
				} else {
					ujid = make([]types.JID, 0)
					jid, _ := types.ParseJID(*m.QuotedMsg.Participant)
					ujid = append(ujid, jid)
				}

			} else if len(m.Querry) > 0 {
				ajid := strings.Split(strings.Trim(m.Querry, " "), ",")
				ujid = make([]types.JID, len(ajid))
				for i, a := range ajid {
					num, err := phonenumbers.Parse(a, "ID")
					if err != nil {
						return
					}
					num_formatted := phonenumbers.Format(num, phonenumbers.E164)
					ujid[i], ok = types.ParseJID(fmt.Sprintf("%s@whatsapp.net", num_formatted[1:]))
					if ok != nil {
						return
					}
				}
			}
			if ujid == nil || len(ujid) == 0 {
				m.Reply("Tag/reply pesan/masukan nomor seseorang yang ingin dikeluarkan dari grup ini.")
				return
			}
			// bisa kick admin btw :v
			// kalau ga lupa kapan² gw buat atau bantu tambahin :v
			resp, err := client.WA.UpdateGroupParticipants(m.From, ujid, whatsmeow.ParticipantChangeAdd)
			if err != nil {
				m.Reply("Mao gagal mengeluarkan user tersebut.")
				return
			}
			for _, item := range resp {
				// bingung
				if item.Error == 409 {
					m.Reply("e")
				} else {
					m.Reply("Mao tidak bisa mengeluarkan user tersebut.")
				}
			}
		},
	})
}
