package main

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/mrd0ll4r/tbotapi"
)

func helpHandler(update tbotapi.Update, bot *tbotapi.TelegramBotAPI) {
	msg := *update.Message

	recipient := tbotapi.NewRecipientFromChat(msg.Chat)

	//markdown
	out := bot.NewOutgoingMessage(recipient, ``)

	out.SetHTML(true)
	out.Text = `
<code>	
package main

import "fmt"

func main() {
    fmt.Println("hello world")
}
</code>
	`

	out.SetReplyKeyboardMarkup(tbotapi.ReplyKeyboardMarkup{
		Keyboard: [][]tbotapi.KeyboardButton{
			[]tbotapi.KeyboardButton{
				tbotapi.KeyboardButton{Text: "Q1: A"},
				tbotapi.KeyboardButton{Text: "Q1: B"},
			}},
		OneTimeKeyboard: true,
	})
	out.Send()
}

func sampleHandler(update tbotapi.Update, bot *tbotapi.TelegramBotAPI) {
	msg := *update.Message

	recipient := tbotapi.NewRecipientFromChat(msg.Chat)

	//markdown
	out := bot.NewOutgoingMessage(recipient, `*bold text* _italic text_ [text](http://google.com)`)
	out.SetMarkdown(true)
	out.Send()

	out.Text = "```func main() { //comment }```"
	out.Send()

	//HTML
	out.SetMarkdown(false) //disable if you want to use HTML or default
	out.SetHTML(true)
	out.Text = `
<b>bold</b>, <strong>bold</strong>
<i>italic</i>, <em>italic</em>
<a href="http://telegram.org">inline http://telegram.org</a>
<code>inline fixed-width code</code>
<pre>pre-formatted fixed-width code block</pre>
	`

	hideKeyboard(out)
	out.Send()
}

func hideKeyboard(out *tbotapi.OutgoingMessage) {
	out.SetReplyKeyboardHide(tbotapi.ReplyKeyboardHide{HideKeyboard: true})
}

func endHandler(update tbotapi.Update, bot *tbotapi.TelegramBotAPI) {
	msg := *update.Message

	recipient := tbotapi.NewRecipientFromChat(msg.Chat)
	out := bot.NewOutgoingMessage(recipient, "End of quiz")

	hideKeyboard(out)
	out.Send()
}

func echoHandler(update tbotapi.Update, bot *tbotapi.TelegramBotAPI) {
	msg := *update.Message

	recipient := tbotapi.NewRecipientFromChat(msg.Chat)
	out := bot.NewOutgoingMessage(recipient, *msg.Text)

	log.Println(*update.Message.Chat.LastName)
	log.Println(*update.Message.Chat.Username)
	log.Println(update.Message.Chat.ID)
	log.Println(update.Message.Chat.String())

	hideKeyboard(out)
	out.Send()
}

func route(update tbotapi.Update, bot *tbotapi.TelegramBotAPI) {
	msg := *update.Message

	txt := *msg.Text

	slc := strings.Split(txt, " ")

	firstword := slc[0]

	if firstword[0:1] == "Q" {
		helpHandler(update, bot)
	} else if firstword[0:1] == "~" {
		endHandler(update, bot)
	} else {

		switch strings.ToUpper(firstword) {
		case "HELP":
			helpHandler(update, bot)

		case "SAMPLE":
			sampleHandler(update, bot)

		default:
			echoHandler(update, bot)

		}
	}
}

func main() {
	//http://telegram.me/gowizardbot

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	bot, err := tbotapi.New("API_TOKEN_HERE")

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
