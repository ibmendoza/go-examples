package main

import (
	"log"

	//api "gopkg.in/telegram-bot-api.v4"
	api "github.com/go-telegram-bot-api/telegram-bot-api"
)

func echoHandler(update api.Update, bot *api.BotAPI) {
	chatID := update.Message.Chat.ID
	msgID := update.Message.MessageID

	msg := api.NewMessage(chatID, "")
	msg.ReplyToMessageID = msgID
	msg.Text = "echo " + update.Message.Text
	bot.Send(msg)
}

func docHandler(update api.Update, bot *api.BotAPI) {
	chatID := update.Message.Chat.ID
	msgID := update.Message.MessageID

	msgdoc := api.NewDocumentUpload(chatID, "C:/gopher.jpg")
	msgdoc.ReplyToMessageID = msgID
	bot.Send(msgdoc)
}

func route(update api.Update, bot *api.BotAPI) {

	//NOTE (a crude event handler)
	//the routing logic is completely arbitrary and it is up to you
	//(like parsing the text or some business logic)

	//replace below to build your own custom route logic

	i := 1

	switch i {
	case 1:
		echoHandler(update, bot)

	case 2:
		docHandler(update, bot)
	case 3:
		log.Println("three")
	default:
	}
}

func main() {

	bot, err := api.NewBotAPI("yourtoken")
	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := api.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		log.Println("From")
		log.Println(update.Message.From.UserName)
		log.Println("Text")
		log.Println(update.Message.Text)
		log.Println("ChatID")
		log.Println(update.Message.Chat.ID)
		log.Println("MessageID")
		log.Println(update.Message.MessageID)
		log.Println("Caption")
		log.Println(update.Message.Caption)
		log.Println("Command")
		log.Println(update.Message.Command())
		log.Println("Date")
		log.Println(update.Message.Date)

		route(update, bot)
	}
}
