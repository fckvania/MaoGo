package libs

import (
	"mao/src/helpers"
	"time"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
)

type IHandler struct {
	Container *store.Device
}

func NewHandler(container *sqlstore.Container) *IHandler {
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}
	return &IHandler{
		Container: deviceStore,
	}
}

func (h *IHandler) Client(jbot ...bool) *whatsmeow.Client {
	clientLog := waLog.Stdout("lient", "ERROR", true)
	client := whatsmeow.NewClient(h.Container, clientLog)
	client.AddEventHandler(h.RegisterHandler(client, jbot...))
	return client
}

func (h *IHandler) RegisterHandler(client *whatsmeow.Client, jbot ...bool) func(evt interface{}) {
	return func(evt interface{}) {
		sock := NewClient(client)
		switch v := evt.(type) {
		case *events.Message:
			m := NewSmsg(v, sock, jbot...)
			if !helpers.Public && !m.IsOwner {
				return
			}
			// Read message
			sock.WA.MarkRead([]string{m.StanzaId}, time.Now(), m.From, m.Sender)
			// Get command
			go Get(sock, m)
			return
		case *events.Connected, *events.PushNameSetting:
			if len(client.Store.PushName) == 0 {
				return
			}
			client.SendPresence(types.PresenceAvailable)
		}
	}
}
