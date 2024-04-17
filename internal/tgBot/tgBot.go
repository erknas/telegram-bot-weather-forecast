package tgBot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/zeze322/telegram-bot-weather-forecast/internal/weather"
	"os"
)

const hPaToMm = 0.75

func BotInit() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("failed to load .env file")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		return fmt.Errorf("failed to init telegram bot")
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		w, err := weather.Forecast(update.Message.Text, os.Getenv("WEATHER_TOKEN"))
		if err != nil {
			continue
		}

		var text string

		if w.Condition == nil {
			text = fmt.Sprintf("Не удалось получить данные о погоде для города `%s`. Ошибка в названии города", update.Message.Text)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			bot.Send(msg)
			continue
		}

		switch w.Condition[0].Description {
		case "пасмурно":
			text = fmt.Sprintf(`%s
Температура: %.f°C, %s  🌥️
Давление: %.f мм рт. ст.
Ветер: %.f м/с
Влажность: %d%%`, w.Name, w.Main.Temp, w.Condition[0].Description, float64(w.Main.Pressure)*hPaToMm, w.Wind.Speed, w.Main.Humidity)
		case "облачно с прояснениями":
			text = fmt.Sprintf(`%s
Температура: %.f°C, %s  ⛅
Давление: %.f мм рт. ст.
Ветер: %.f м/с
Влажность: %d%%`, w.Name, w.Main.Temp, w.Condition[0].Description, float64(w.Main.Pressure)*hPaToMm, w.Wind.Speed, w.Main.Humidity)
		case "ясно":
			text = fmt.Sprintf(`%s
Температура: %.f°C, %s  ☀️
Давление: %.f мм рт. ст.
Ветер: %.f м/с
Влажность: %d%%`, w.Name, w.Main.Temp, w.Condition[0].Description, float64(w.Main.Pressure)*hPaToMm, w.Wind.Speed, w.Main.Humidity)
		case "плотный туман":
			text = fmt.Sprintf(`%s
Температура: %.f°C, %s  🌫️
Давление: %.f мм рт. ст.
Ветер: %.f м/с
Влажность: %d%%`, w.Name, w.Main.Temp, w.Condition[0].Description, float64(w.Main.Pressure)*hPaToMm, w.Wind.Speed, w.Main.Humidity)
		case "небольшой снег":
			text = fmt.Sprintf(`%s
Температура: %.f°C, %s  🌨️
Давление: %.f мм рт. ст.️
Ветер: %.f м/с
Влажность: %d%%`, w.Name, w.Main.Temp, w.Condition[0].Description, float64(w.Main.Pressure)*hPaToMm, w.Wind.Speed, w.Main.Humidity)
		case "дождь":
			text = fmt.Sprintf(`%s
Температура: %.f°C, %s  🌧️
Давление: %.f мм рт. ст.
Ветер: %.f м/с
Влажность: %d%%`, w.Name, w.Main.Temp, w.Condition[0].Description, float64(w.Main.Pressure)*hPaToMm, w.Wind.Speed, w.Main.Humidity)
		case "переменная облачность":
			text = fmt.Sprintf(`%s
Температура: %.f°C, %s  ☁️
Давление: %.f мм рт. ст.
Ветер: %.f м/с
Влажность: %d%%`, w.Name, w.Main.Temp, w.Condition[0].Description, float64(w.Main.Pressure)*hPaToMm, w.Wind.Speed, w.Main.Humidity)
		case "небольшой дождь":
			text = fmt.Sprintf(`%s
Температура: %.f°C, %s  🌨️
Давление: %.f мм рт. ст.
Ветер: %.f м/с
Влажность: %d%%`, w.Name, w.Main.Temp, w.Condition[0].Description, float64(w.Main.Pressure)*hPaToMm, w.Wind.Speed, w.Main.Humidity)
		case "небольшая облачность":
			text = fmt.Sprintf(`%s
Температура: %.f°C, %s  ⛅
Давление: %.f мм рт. ст.
Ветер: %.f м/с
Влажность: %d%%`, w.Name, w.Main.Temp, w.Condition[0].Description, float64(w.Main.Pressure)*hPaToMm, w.Wind.Speed, w.Main.Humidity)
		case "сильный дождь":
			text = fmt.Sprintf(`%s
Температура: %.f°C, %s  🌨️
Давление: %.f мм рт. ст.
Ветер: %.f м/с
Влажность: %d%%`, w.Name, w.Main.Temp, w.Condition[0].Description, float64(w.Main.Pressure)*hPaToMm, w.Wind.Speed, w.Main.Humidity)
		default:
			text = fmt.Sprintf(`%s
Температура: %.f°C, %s
Давление: %.f мм рт. ст.
Ветер: %.f м/с
Влажность: %d%%`, w.Name, w.Main.Temp, w.Condition[0].Description, float64(w.Main.Pressure)*hPaToMm, w.Wind.Speed, w.Main.Humidity)
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		bot.Send(msg)
	}

	return nil
}
