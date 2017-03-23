package main

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/mrd0ll4r/tbotapi"
)

var m map[int]string

func initQuestion() {
	counter = 0

	m = make(map[int]string)
	m[1] = "What is the capital of the Philippines?"
	m[2] = "What is the capital of the USA?"
	m[3] = "What is the capital of Japan?"
	m[4] = "What is the capital of Canada?"
	m[5] = "What is the capital of Russia?"
}

func getQuestion(update tbotapi.Update) (string, int) {
	/*
	   	return `
	   <code>
	   package main
	   import "fmt"
	   func main() {
	       fmt.Println("hello world")
	   }
	   </code>
	   	`
	*/
	if m == nil {
		return "", counter
	} else {
		return m[counter], counter
	}
}

func startHandler(update tbotapi.Update, bot *tbotapi.TelegramBotAPI) {
	initQuestion()

	msg := *update.Message
	recipient := tbotapi.NewRecipientFromChat(msg.Chat)

	out := bot.NewOutgoingMessage(recipient, ``)
	out.SetHTML(true)
	out.Text = "Get ready..."

	hideKeyboard(out)

	out.Send()

	questionHandler(update, bot)
}

var counter = 0

//control logic
func isEoQ(update tbotapi.Update) bool {
	return counter == 6
}

func (se staticExample) retrieveChoices(id int) (c1, c2, c3, c4 string) {
	switch id {
	case 1:
		c1 = "Manila"
		c2 = "Makati"
		c3 = "Quezon City"
		c4 = "Mall of Asia"

	case 2:
		c1 = "Washington, D.C."
		c2 = "Washington state"
		c3 = "Los Angeles, CA"
		c4 = "New York, NY"

	case 3:
		c1 = "Osaka"
		c2 = "Hokkaido"
		c3 = "Tokyo"
		c4 = "Ibaraki"

	case 4:
		c1 = "Vancouver"
		c2 = "Ontario"
		c3 = "Ottawa"
		c4 = "Toronto"

	case 5:
		c1 = "Vladivostok"
		c2 = "Putin"
		c3 = "Saint Petersburg"
		c4 = "Moscow"
	}

	return
}

type iRetrieveChoices interface {
	retrieveChoices(id int) (c1, c2, c3, c4 string)
}

type staticExample struct{}

//proxy function
func getChoices(id int) (c1, c2, c3, c4 string) {
	se := new(staticExample)
	c1, c2, c3, c4 = se.retrieveChoices(id)

	return
}

func questionHandler(update tbotapi.Update, bot *tbotapi.TelegramBotAPI) {
	counter++ //simulate state

	msg := *update.Message

	recipient := tbotapi.NewRecipientFromChat(msg.Chat)

	//markdown
	out := bot.NewOutgoingMessage(recipient, ``)

	out.SetHTML(true)

	if isEoQ(update) {
		out.Text = "End of questions"
		hideKeyboard(out)
	} else {
		var id int
		out.Text, id = getQuestion(update)

		c1, c2, c3, c4 := getChoices(id)

		setKeyboard(out, c1, c2, c3, c4)
		//setKeyboard2(out, c1, c2, c3, c4)
	}

	//logUserDetails(update)

	out.Send()

	m := update.Message
	typ := m.Type()
	if typ != tbotapi.TextMessage {
		// Ignore non-text messages for now.
		log.Println("Ignoring non-text message")
		return
	}

	log.Println(*m.Text) //get the response
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

func setKeyboard(out *tbotapi.OutgoingMessage, str1, str2, str3, str4 string) {
	out.SetReplyKeyboardMarkup(tbotapi.ReplyKeyboardMarkup{
		Keyboard: [][]tbotapi.KeyboardButton{
			[]tbotapi.KeyboardButton{
				tbotapi.KeyboardButton{Text: str1},
				tbotapi.KeyboardButton{Text: str2},
				tbotapi.KeyboardButton{Text: str3},
				tbotapi.KeyboardButton{Text: str4},
			}},
		OneTimeKeyboard: true,
		ResizeKeyboard:  true,
	})
}

func setKeyboard2(out *tbotapi.OutgoingMessage, str1, str2, str3, str4 string) {
	out.SetReplyKeyboardMarkup(tbotapi.ReplyKeyboardMarkup{
		Keyboard: [][]tbotapi.KeyboardButton{
			[]tbotapi.KeyboardButton{
				tbotapi.KeyboardButton{Text: str1},
				tbotapi.KeyboardButton{Text: str2},
				tbotapi.KeyboardButton{Text: str3},
				tbotapi.KeyboardButton{Text: str4},
				tbotapi.KeyboardButton{Text: "5"},
				tbotapi.KeyboardButton{Text: "6"},
			}},
		OneTimeKeyboard: true,
	})
}

func hideKeyboard(out *tbotapi.OutgoingMessage) {
	out.SetReplyKeyboardHide(tbotapi.ReplyKeyboardHide{HideKeyboard: true})
}

func logUserDetails(update tbotapi.Update) {
	log.Println(*update.Message.Chat.FirstName)
	log.Println(*update.Message.Chat.LastName)
	log.Println(*update.Message.Chat.Username)
	log.Println(update.Message.Chat.ID)
	log.Println(update.Message.Chat.String())
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

	//logUserDetails(update)

	hideKeyboard(out)
	out.Send()
}

//parse the message, then route
func route(update tbotapi.Update, bot *tbotapi.TelegramBotAPI) {
	msg := *update.Message

	txt := *msg.Text

	//represents the response from multiple choice
	slc := strings.Split(txt, " ")

	firstword := slc[0]

	switch strings.ToUpper(firstword) {
	case "START":
		startHandler(update, bot)

	case "SAMPLE":
		sampleHandler(update, bot)

	case "ECHO":
		echoHandler(update, bot)

	default:
		questionHandler(update, bot)

	}
}

func main() {
	//http://telegram.me/gowizardbot

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	//bot, err := tbotapi.New("API_TOKEN")

	bot, err := tbotapi.New("346080776:AAFa3c2MO5j1AOnj8UHd6YokWgu48nCIcac")
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
