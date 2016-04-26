package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"bitbucket.org/mrd0ll4r/tbotapi"
)

func echoHandler(update tbotapi.Update, bot *tbotapi.TelegramBotAPI) {
	msg := *update.Message

	log.Println(*msg.Text)

	bot.NewOutgoingMessage(tbotapi.NewRecipientFromChat(msg.Chat), *msg.Text).Send()
}

func route(update tbotapi.Update, bot *tbotapi.TelegramBotAPI) {
	i := 1

	switch i {
	case 1:
		echoHandler(update, bot)

	case 2:
		//docHandler(update, bot)

	default:
	}
}

func main() {
	
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	bot, err := tbotapi.New("YourToken")

	if err != nil {
		log.Fatal(err)
	}

	defer bot.Close()

	log.Println("User ID: ", bot.ID)
	log.Println("Bot Name: ", bot.Name)
	log.Println("Bot Username: ", bot.Username)

	for {
		select {
		case <-sigChan:
			signal.Stop(sigChan)
			bot.Close()
			close(sigChan)
			os.Exit(1)

		case update := <-bot.Updates:
			go route(update.Update(), bot)
		}
	}
}
