package commands

import (
	"fmt"
	"mao/src/libs"
	"strings"
)

type item struct {
	Name     string
	IsPrefix bool
}

func menu(client *libs.NewClientImpl, m *libs.IMessage) {
	var str string
	str += fmt.Sprintf("Hello %s, Berikut List Command Yang Tersedia\n\n", m.PushName)
	var tags map[string][]item
	for _, list := range libs.GetList() {
		if tags == nil {
			tags = make(map[string][]item)
		}
		if _, ok := tags[list.Tags]; !ok {
			tags[list.Tags] = []item{}
		}
		tags[list.Tags] = append(tags[list.Tags], item{Name: list.Name, IsPrefix: list.IsPrefix})
	}
	for key := range tags {
		str += fmt.Sprintf(" *%s*\n", strings.ToUpper(key))
		for _, e := range tags[key] {
			var prefix string
			if e.IsPrefix {
				prefix = "#"
			} else {
				prefix = ""
			}
			str += fmt.Sprintf("ゝ %s%s\n", prefix, e.Name)
		}
		str += "\n"
	}
	client.SendWithNewsLestter(m.From, str, "120363205826017795@newsletter", 100, "My Name Mao", m.ContextInfo)
}

func init() {
	libs.NewCommands(&libs.ICommand{
		Name:     "menu",
		Tags:     "main",
		IsPrefix: true,
		Exec:     menu,
	})
}
