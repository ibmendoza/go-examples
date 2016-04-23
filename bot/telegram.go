//https://godoc.org/gopkg.in/telegram-bot-api.v4

package main

import (
	"log"

	api "gopkg.in/telegram-bot-api.v4"
)

func main() {

	bot, err := api.NewBotAPI("yourtokenhere")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := api.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		chatID := update.Message.Chat.ID
		msgID := update.Message.MessageID

		msg := api.NewMessage(chatID, update.Message.Text)
		msg.ReplyToMessageID = msgID
		msg.Text = "echo " + update.Message.Text
		bot.Send(msg)

		//msgdoc := api.NewDocumentUpload(chatID, "C:/gopher.jpg")
		//msgdoc.ReplyToMessageID = msgID
		//bot.Send(msgdoc)
	}
}
