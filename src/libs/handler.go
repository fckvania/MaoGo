package libs

import (
	"mao/src/helpers"

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

func (h *IHandler) Client() *whatsmeow.Client {
	clientLog := waLog.Stdout("lient", "ERROR", true)
	client := whatsmeow.NewClient(h.Container, clientLog)
	client.AddEventHandler(h.RegisterHandler(client))
	return client
}

func (h *IHandler) RegisterHandler(client *whatsmeow.Client) func(evt interface{}) {
	return func(evt interface{}) {
		sock := NewClient(client)
		switch v := evt.(type) {
		case *events.Message:
			m := NewSmsg(v, sock)
			if !helpers.Public && !m.IsOwner {
				return
			}
			go Get(sock, m)
			return
		case *events.LoggedOut:
			con := evt.(*events.LoggedOut)
			if !con.OnConnect {
				return
			}
			break
		case *events.Connected, *events.PushNameSetting:
			if len(client.Store.PushName) == 0 {
				return
			}
			client.SendPresence(types.PresenceAvailable)
		}
	}
}
