package main

import (
	"gobotapi"
	"gobotapi/filters"
	"gobotapi/methods"
	"gobotapi/types"
	"os"
)

func main() {
	songId := "3vPXCd0uNm66sbDEDG5Vwd"
	Download(songId, "m8nstar - thasup")
	token := os.Getenv("TOKEN")
	client := gobotapi.NewClient(token)

	client.OnAnyMessageEvent(filters.Filter(func(client *gobotapi.Client, update types.Message) {
		_, err := client.Invoke(&methods.SendMessage{
			ChatID: update.Chat.ID,
			Text:   "miao",
			ReplyMarkup: &types.InlineKeyboardMarkup{
				InlineKeyboard: [][]types.InlineKeyboardButton{
					{
						{
							Text: "Click me!",
							URL:  "https://branchscope.com",
						},
					},
				},
			},
		})
		if err != nil {
			panic(err)
		}
	}, filters.Command("start")))

	client.OnAnyMessageEvent(filters.Filter(func(client *gobotapi.Client, update types.Message) {
		_, err := client.Invoke(&methods.SendAudio{
			ChatID: update.Chat.ID,
			Audio:  types.InputFile(songId + ".mp3"),
		})
		if err != nil {
			panic(err)
		}
	}, filters.Command("m8nstar")))

	_ = client.Run()
}
