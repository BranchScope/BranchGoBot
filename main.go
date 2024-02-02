package main

import (
	"gobotapi"
	"gobotapi/filters"
	"gobotapi/methods"
	"gobotapi/types"
)

func main() {
	client := gobotapi.NewClient("")

	client.OnAnyMessageEvent(filters.Filter(func(client *gobotapi.Client, update types.Message) {
		_, err := client.Invoke(&methods.SendMessage{
			ChatID: update.Chat.ID,
			Text:   "Helo :D",
		})
		if err != nil {
			panic(err)
		}
	}, filters.Command("start")))

	_ = client.Run()
}
