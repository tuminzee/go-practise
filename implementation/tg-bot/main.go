package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"
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
	log.Println("Bot is starting...")
	log.Println(b.Me)

	log.Println("Bot is now running. Press Ctrl-C to exit.")
	b.Start()

}
