package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/zeze322/telegram-bot-weather-forecast/internal/weather"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		w, err := weather.Forecast(update.Message.Text, os.Getenv("WEATHER_TOKEN"))
		if err != nil {
			log.Fatal(err)
		}

		var text string

		switch w.Condition[0].Description {
		case "пасмурно":
			text = fmt.Sprintf(`%s:
Температура %.f°C, %s  🌥️
Давление %dмм рт. ст.
Ветер %.f м/с`, w.Name, w.Main.Temp, w.Condition[0].Description, w.Main.Pressure, w.Wind.Speed)
		case "облачно с прояснениями":
			text = fmt.Sprintf(`%s:
Температура %.f°C, %s  ⛅
Давление %dмм рт. ст.
Ветер %.f м/с`, w.Name, w.Main.Temp, w.Condition[0].Description, w.Main.Pressure, w.Wind.Speed)
		case "ясно":
			text = fmt.Sprintf(`%s:
Температура %.f°C, %s  ☀️
Давление %dмм рт. ст.
Ветер %.f м/с`, w.Name, w.Main.Temp, w.Condition[0].Description, w.Main.Pressure, w.Wind.Speed)
		case "плотный туман":
			text = fmt.Sprintf(`%s:
Температура %.f°C, %s  🌫️
Давление %dмм рт. ст.
Ветер %.f м/с`, w.Name, w.Main.Temp, w.Condition[0].Description, w.Main.Pressure, w.Wind.Speed)
		case "небольшой снег":
			text = fmt.Sprintf(`%s:
Температура %.f°C, %s  ❄️
Давление %dмм рт. ст.️
Ветер %.f м/с`, w.Name, w.Main.Temp, w.Condition[0].Description, w.Main.Pressure, w.Wind.Speed)
		case "дождь":
			text = fmt.Sprintf(`%s:
Температура %.f°C, %s  🌧️
Давление %dмм рт. ст.
Ветер %.f м/с`, w.Name, w.Main.Temp, w.Condition[0].Description, w.Main.Pressure, w.Wind.Speed)
		case "переменная облачность":
			text = fmt.Sprintf(`%s:
Температура %.f°C, %s  ☁️
Давление %dмм рт. ст.
Ветер %.f м/с`, w.Name, w.Main.Temp, w.Condition[0].Description, w.Main.Pressure, w.Wind.Speed)
		case "небольшой дождь":
			text = fmt.Sprintf(`%s:
Температура %.f°C, %s  🌧️
Давление %dмм рт. ст.
Ветер %.f м/с`, w.Name, w.Main.Temp, w.Condition[0].Description, w.Main.Pressure, w.Wind.Speed)
		default:
			text = fmt.Sprintf(`%s:
Температура %.f°C, %s
Давление %dмм рт. ст.
Ветер %.f м/с`, w.Name, w.Main.Temp, w.Condition[0].Description, w.Main.Pressure, w.Wind.Speed)
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		bot.Send(msg)
	}
}
