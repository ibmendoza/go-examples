package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
	"github.com/johnnylee/sqlxchain"
	"github.com/mrd0ll4r/tbotapi"
)

var m map[int]string
var n map[string]string
var dbx *sqlxchain.SqlxChain
var db *sqlx.DB

//websocketd --port=8080 --staticdir=. survey

func initQuestion() {
	//counter = 0

	n = make(map[string]string)
	n["a:1"] = "In a scale of 1 to 4, was the agent polite?"
	n["a:2"] = "Calls were able to get through to a ringing tone and answered within 15 seconds"
	n["a:3"] = `
		I felt the officer made efforts to understand my problems, regardless of the outcome. 	
	`
	n["a:4"] = "What is the capital of Canada?"
	n["a:5"] = "What is the capital of Russia?"
	n["a:6"] = "Thank you for taking the survey"
}

func startHandler(update tbotapi.Update, bot *tbotapi.TelegramBotAPI) {
	initQuestion()
	//initHandler(update, bot)
}

func retrieveChoices(id string) (c1, c2, c3, c4, c5, c6 string) {
	switch id {
	case "a:1":
		c1 = "#a:1   one"
		c2 = "#a:1   two"
		c3 = "#a:1   three"
		c4 = "#a:1   four"

	case "a:2":
		c1 = "#a:2   Yes"
		c2 = "#a:2   No"

	case "a:3":
		c1 = "#a:3   Agree"
		c2 = "#a:3   Strongly agree"
		c3 = "#a:3   Disagree"
		c4 = "#a:3   Strongly disagree"

	case "a:4":
		c1 = "#a:4 Vancouver"
		c2 = "#a:4 Ontario"
		c3 = "#a:4 Ottawa"
		c4 = "#a:4 Toronto"

	case "a:5":
		c1 = "#a:5 Vladivostok"
		c2 = "#a:5 Putin"
		c3 = "#a:5 Saint Petersburg"
		c4 = "#a:5 Moscow"
	}

	return
}

func alert(args ...interface{}) {
	log.Println(args)
}

/*
func initHandler(update tbotapi.Update, bot *tbotapi.TelegramBotAPI) {
	msg := *update.Message
	response := *update.Message.Text

	alert("response, ", response)

	m := update.Message
	typ := m.Type()
	if typ != tbotapi.TextMessage {
		// Ignore non-text messages for now.
		log.Println("Ignoring non-text message")
		return
	}
	slc := strings.Split(response, " ")
	firstword := slc[0]
	question := firstword[1:len(firstword)]
	slcAnswer := strings.Split(response, question)

	if len(slcAnswer) != 2 {
		return
	}
	//alert("response ", response)
	//alert("question, ", question)
	//alert("answer ", slcAnswer[1])
	recipient := tbotapi.NewRecipientFromChat(msg.Chat)

	out := bot.NewOutgoingMessage(recipient, ``)
	out.SetHTML(true)
	out.Text = n["a:1"]
	c1 := "#a:1 Manila"
	c2 := "#a:1 Makati"
	c3 := "#a:1 Quezon City"
	c4 := "#a:1 Mall of Asia"
	setKeyboard(out, c1, c2, c3, c4, c5, c6)

	out.Send()
}
*/

func questionHandler(update tbotapi.Update, bot *tbotapi.TelegramBotAPI) {
	msg := *update.Message
	response := *update.Message.Text

	m := update.Message
	typ := m.Type()
	if typ != tbotapi.TextMessage {
		// Ignore non-text messages for now.
		log.Println("Ignoring non-text message")
		return
	}
	slc := strings.Split(response, " ")

	firstword := slc[0] //#a:1

	question := firstword[1:len(firstword)]
	alert("question: ", question)

	//index is the short name for question index
	index := firstword[1:len(firstword)] //a:1

	//alert("question, ", index)

	//turn a:1 to a:2
	slcIndex := strings.Split(index, ":")
	idx := slcIndex[1] //from a:1 slcIndex[1] equals 1

	//("idx ", idx)

	alert("idxNumber ", idx)
	var newIdx string
	//increment by 1
	i, err := strconv.Atoi(idx)
	if err != nil {
		return
	}
	i++
	newIdx = strconv.Itoa(i)

	slcAnswer := strings.Split(response, index)
	alert("slcAnswer ", slcAnswer)
	if len(slcAnswer) != 2 {
		return
	}
	answer := slcAnswer[1]
	alert("answer ", answer)

	//insert to db
	sql := `
		insert into metrics(questionidx, question, answer)
		values(?, ?, ?)	
	`

	//var id int64
	//insert to database
	ans := strings.Trim(answer, " ")
	err = dbx.Context().Begin().
		Exec(sql, question, " ", ans).
		Commit().
		//LastInsertId(&id).
		Err()

	if err != nil {
		alert("Error: ", err)
	}

	//display to stdout for pickup by websocketd
	fmt.Println("question: ", question, " answer: ", ans)

	recipient := tbotapi.NewRecipientFromChat(msg.Chat)

	out := bot.NewOutgoingMessage(recipient, ``)
	out.SetHTML(true)

	//should be out.Text = n["a:2"]
	out.Text = n["a:"+newIdx]

	c1, c2, c3, c4, c5, c6 := retrieveChoices("a:" + newIdx)
	setKeyboard(out, c1, c2, c3, c4, c5, c6)

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

func setKeyboard(out *tbotapi.OutgoingMessage, str1, str2, str3, str4, str5, str6 string) {
	out.SetReplyKeyboardMarkup(tbotapi.ReplyKeyboardMarkup{
		Keyboard: [][]tbotapi.KeyboardButton{
			[]tbotapi.KeyboardButton{
				tbotapi.KeyboardButton{Text: str1},
				tbotapi.KeyboardButton{Text: str2},
				tbotapi.KeyboardButton{Text: str3},
				tbotapi.KeyboardButton{Text: str4},
				tbotapi.KeyboardButton{Text: str5},
				tbotapi.KeyboardButton{Text: str6},
			}},
		OneTimeKeyboard: true,
		ResizeKeyboard:  true,
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
	out := bot.NewOutgoingMessage(recipient, "End of survey")

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

	//alert("firstword ", firstword)

	//example firstword: #a:1 Manila
	if string(firstword[0]) == "#" {
		questionHandler(update, bot)
	} else {

		switch strings.ToUpper(firstword) {
		/*
			case "START":
				startHandler(update, bot)
		*/

		case "SAMPLE":
			sampleHandler(update, bot)

		case "ECHO":
			echoHandler(update, bot)

		default:
			//initHandler(update, bot)
		}
	}
}

func main() {
	//http://telegram.me/gowizardbot

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	//bot, err := tbotapi.New("API_TOKEN")

	bot, err := tbotapi.New("token_here")
	if err != nil {
		log.Fatal(err)
	}

	defer bot.Close()

	log.Println("User ID: ", bot.ID)
	log.Println("Bot Name: ", bot.Name)
	log.Println("Bot Username: ", bot.Username)

	//connect to database
	server := "localhost" //"RMAJSNB0023"
	user := `SAuser`      //"systemadmin"
	pswd := "RM@hrmsg0!"  //"8rm@syst3m8"

	dsn := "server=" + server + ";user id=" + user + ";password=" + pswd + ";database=dbrma"
	//db, err := sql.Open("mssql", dsn)
	db = sqlx.MustConnect("mssql", dsn)
	dbx, err = sqlxchain.New("mssql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Println("From Open() attempt: " + err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Println("From Ping() Attempt: " + err.Error())
		return
	}
	log.Println("Connected to database...")
	//connect to database

	initQuestion()

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
