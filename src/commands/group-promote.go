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
		Name:     "(promote|pm)",
		As:       []string{"promote"},
		Tags:     "group",
		IsPrefix: true,
		IsGroup:  true,
		IsAdmin:  true,
		Exec: func(client *libs.NewClientImpl, m *libs.IMessage) {
			var ujid []types.JID
			var ok error
			// apalah ini gw bingung
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
				m.Reply("Tag atau balas pesan seseorang yang mau dijadikan admin")
				return
			}
			fmt.Println("All user jid ", ujid)
			_, err := client.WA.UpdateGroupParticipants(m.From, ujid, whatsmeow.ParticipantChangePromote)
			if err != nil {
				m.Reply("Mao gagal menjadikan user tersebut :(\nsilankan cek kembali nomor tersebut")
				return
			}
			// res, _ := client.WA.GetGroupInfo(m.From)
			// for _, item := range resp {
			// 	if item.Error == 403 && item.AddRequest != nil {
			// 		client.WA.SendMessage(context.TODO(), item.JID, &waProto.Message{
			// 			GroupInviteMessage: &waProto.GroupInviteMessage{
			// 				InviteCode:       proto.String(item.AddRequest.Code),
			// 				InviteExpiration: proto.Int64(item.AddRequest.Expiration.Unix()),
			// 				GroupJid:         proto.String(m.From.String()),
			// 				GroupName:        proto.String(res.Name),
			// 				Caption:          proto.String(""),
			// 			},
			// 		})

			// 	} else if item.Error == 409 {
			// 		m.Reply("nomor tersebut sudah ada di group ini kak")
			// 	} else {
			// 		m.Reply("Mao tidak bisa menambahkan user tersebut :)")
			// 	}
			// }
		},
	})
}
