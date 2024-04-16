package main

import (
	"github.com/zeze322/telegram-bot-weather-forecast/internal/tgBot"
	"log"
)

func main() {
	if err := tgBot.BotInit(); err != nil {
		log.Fatal(err)
	}
	
}
