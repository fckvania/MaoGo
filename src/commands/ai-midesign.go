package commands

import (
	"mao/src/libs"
	"mao/src/libs/api"
)

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:     "(miscrosoftdesigner|midesign)",
		As:       []string{"midesign"},
		Tags:     "ai",
		IsPrefix: true,
		IsQuerry: true,
		IsWaitt:  true,
		Exec: func(client *libs.NewClientImpl, m *libs.IMessage) {
			data, err := api.MicrosoftDesigner(m.Querry)
			if err != nil {
				m.Reply(err.Error())
				return
			}

			buffer := client.GetBytes(data["image_urls_thumbnail"].([]interface{})[0].(map[string]interface{})["ImageUrl"].(string))

			client.SendImage(m.From, buffer, m.Querry, m.ID)
		},
	})
}
