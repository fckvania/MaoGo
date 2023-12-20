package commands

import (
	"context"
	"fmt"
	"mao/src/libs"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/skip2/go-qrcode"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
)

type IJadibot struct {
	Client *whatsmeow.Client
	Number string
	Type   string
}

var queque = make(map[string]bool)
var jRoom = make(map[string]IJadibot)

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:      "(jadibot|stopjadibot)",
		As:        []string{"jadibot", "stopjadibot"},
		Tags:      "main",
		IsPrefix:  true,
		IsPrivate: true,
		After: func(client *libs.NewClientImpl, m *libs.IMessage) {
			if queque[m.Sender.ToNonAD().String()] && strings.Contains(m.QuotedMsg.GetStanzaId(), "JBOT") {
				pattern := regexp.MustCompile(`[1-2]`)
				delete(queque, m.Sender.ToNonAD().String())
				if pattern.MatchString(m.Body) {
					var isConnect bool
					os.Mkdir(".sesi", 0777)
					sesiPath := fmt.Sprintf(".sesi/%s.db", m.Sender.ToNonAD().String())
					dbLog := waLog.Stdout("Database", "ERROR", true)
					container, err := sqlstore.New("sqlite3", fmt.Sprintf("file:%s?_foreign_keys=on", sesiPath), dbLog)
					if err != nil {
						panic(err)
					}
					handler := libs.NewHandler(container)
					conn := handler.Client()
					conn.AddEventHandler(func(evt interface{}) {
						switch evt.(type) {
						case *events.LoggedOut:
							con := evt.(*events.LoggedOut)
							if !con.OnConnect {
								client.SendText(m.From, "Logout succes", nil)
								os.Remove(sesiPath)
								return
							}
							break
						}
					})
					conn.PrePairCallback = func(jid types.JID, platform, businessName string) bool {
						m.Reply("Success Login ID :\n\nNumber: " + jid.User + "\nPlatform: " + platform + "\nBusinessName: " + businessName)
						isConnect = true
						jRoom[m.Sender.ToNonAD().String()] = IJadibot{
							Client: conn,
							Number: jid.User,
						}
						return true
					}
					conn.Disconnect()

					if conn.Store.ID == nil {
						switch string(m.Body) {
						case "1":
							if err := conn.Connect(); err != nil {
								panic(err)
							}

							code, err := conn.PairPhone(m.Sender.User, true, whatsmeow.PairClientChrome, "Chrome (Linux)")
							if err != nil {
								panic(err)
							}

							res, _ := m.Reply("Code Kamu : " + code + "\n\nExpired : 2 menit")
							go func() {
								for range time.Tick(2 * time.Minute) {
									if isConnect {
										return
									}
									conn.Disconnect()
									conn.Logout()
									client.DeleteMsg(m.From, res.ID, true)
									m.Reply("Expired")
									os.Remove(sesiPath)
									delete(jRoom, m.Sender.ToNonAD().String())
									break
								}
							}()

							jRoom[m.Sender.ToNonAD().String()] = IJadibot{
								Type:   "Pairing",
								Number: m.Sender.User,
								Client: nil,
							}
							break
						case "2":
							qrChan, _ := conn.GetQRChannel(context.Background())
							if err := conn.Connect(); err != nil {
								panic(err)
							}
							var i int = 1
							var limit = 3
							for evt := range qrChan {
								if evt.Event == "code" {
									if i > 3 {
										conn.Disconnect()
										conn.Logout()
										m.Reply("Limit Qr Sudah Habis")
										os.Remove(sesiPath)
										delete(jRoom, m.Sender.ToNonAD().String())
										break
									}
									var png []byte
									png, _ = qrcode.Encode(evt.Code, qrcode.Medium, 256)
									ms := evt.Timeout.Microseconds() / 1000000
									minutes := ms / 60
									remainingSeconds := ms - minutes*60
									res, _ := client.SendImage(m.From, png, fmt.Sprintf("Expired : %v menit, %v detik\nLimit Sisa : %v/%v\n\nSilahkan Di Scan ya...", minutes, remainingSeconds, limit, 3), nil)
									limit--
									i++
									duration := evt.Timeout / time.Second
									go func() {
										for range time.Tick(duration * time.Second) {
											client.DeleteMsg(m.From, res.ID, true)
											break
										}
									}()
									jRoom[m.Sender.ToNonAD().String()] = IJadibot{
										Type:   "Qr",
										Number: m.Sender.User,
										Client: nil,
									}
								}
							}
						}
					} else {
						conn.Connect()
						m.Reply("Sucess Reconnect")
						jRoom[m.Sender.ToNonAD().String()] = IJadibot{
							Client: conn,
							Number: conn.Store.ID.User,
						}
					}
				}
			}

		},
		Exec: func(client *libs.NewClientImpl, m *libs.IMessage) {
			if reg, _ := regexp.MatchString(`(stopjadibot)`, m.Command); reg {
				if !m.IsBot {
					return
				}

				m.Reply("Success Logout")
				client.WA.Disconnect()
				client.WA.Logout()
				os.Remove(fmt.Sprintf(".sesi/%s.db", m.Sender.ToNonAD().String()))
			} else {
				if m.IsBot {
					m.Reply("Maaf, Bot Tidak Bisa Menggunakan Perintah Ini")
					return
				}

				if jRoom[m.Sender.ToNonAD().String()].Client != nil {
					m.Reply("Maaf, Kamu Sudah Login")
					return
				}

				if queque[m.Sender.ToNonAD().String()] {
					m.Reply("Silahkan Pilih Mode Login Terlebih Dahuulu")
					return
				}

				if jRoom[m.Sender.ToNonAD().String()].Number != "" && jRoom[m.Sender.ToNonAD().String()].Client == nil {
					m.Reply("Maaf, Silahkan Selesaikan Mode " + jRoom[m.Sender.ToNonAD().String()].Type + " Terlebih Dahulu")
					return
				}

				if len(jRoom) > 5 {
					m.Reply("Maaf, Limit Jadi Bot Sudah Habis")
				}

				res, _ := m.Reply("Silahkan Pilih Mode Login :\n\n1. Pairing (Nomor)\n2. Qr\n\n*NT:* Balas Pesan ini, Dan *KETIK NOMORNYA*, Untuk Memilih Ya.", whatsmeow.SendRequestExtra{
					ID: client.GenerateMessageID("JBOT"),
				})

				queque[m.Sender.ToNonAD().String()] = true
				for range time.Tick(3 * time.Minute) {
					if queque[m.Sender.ToNonAD().String()] {
						delete(queque, m.Sender.ToNonAD().String())
						client.DeleteMsg(m.From, res.ID, true)
						break
					}
				}
			}
		},
	})
}
