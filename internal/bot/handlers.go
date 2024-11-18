package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func PingHandler(upd tgbotapi.Update) tgbotapi.Chattable {
	return tgbotapi.NewMessage(upd.FromChat().ID, "pong")
}
