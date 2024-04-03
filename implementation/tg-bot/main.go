package main

import (
	"fmt"
	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	botToken := os.Getenv("BOT_TOKEN")

	pref := tele.Settings{
		Token: botToken,
	}
	b, err := tele.NewBot(pref)

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(c tele.Context) error {
		////return c.Send("Hello world!")
		//inlineBtn := &tele.InlineButton{
		//	WebApp: &tele.WebApp{
		//		URL: "https://predict.wagmi11.com/",
		//	},
		//}
		m := b.NewMarkup()
		if err != nil {
			fmt.Println(err)
		}
		return c.Send(m)
	})

	fmt.Println(b)

	b.Start()

}
