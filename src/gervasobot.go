package main

import (
	"log"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var bot *tb.Bot

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	bot = b

	if err != nil {
		log.Fatal(err)
		return
	}

	bot.Handle("/digievolvi", update)
	bot.Start()
}

func update(m *tb.Message) {
	log.Printf(m.Payload)
	if isAuthorized(m) {
		response := execPullCommand()
		bot.Send(m.Sender, response)
		response = buildProgram()
		bot.Send(m.Sender, "Tutto ok")
	} else {
		bot.Send(m.Sender, "Mi spiace "+m.Sender.Username+" non sei autorizzato")
	}
}

func isAuthorized(m *tb.Message) bool {
	return m.Payload == os.Getenv("MY_SECRET_PASSWORD") && m.Sender.Username == os.Getenv("MY_AUTH_USER")
}
