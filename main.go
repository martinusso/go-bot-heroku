package main

import (
	"os"

	"github.com/go-chat-bot/bot/rocket"
	_ "github.com/go-chat-bot/plugins/catfacts"
	_ "github.com/go-chat-bot/plugins/catgif"
	_ "github.com/go-chat-bot/plugins/chucknorris"
	_ "github.com/go-chat-bot/plugins/godoc"
)

func main() {
	config := &rocket.Config{
		Server:   os.Getenv("ROCKET_SERVER"),
		Port:     os.Getenv("ROCKET_PORT"),
		User:     os.Getenv("ROCKET_USER"),
		Email:    os.Getenv("ROCKET_EMAIL"),
		Password: os.Getenv("ROCKET_PASSWORD"),
		UseTLS:   false,
		Debug:    os.Getenv("DEBUG") != "",
	}
	rocket.Run(config)
}
