package bot

import (
	"context"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func RunHandler(upd tgbotapi.Update, ch chan tgbotapi.Chattable) {
	if msg := upd.Message; msg.Text == "ping" {
		ch <- PingHandler(upd)
	}
}

func (b *Bot) HandleUpdate(ctx context.Context, upd tgbotapi.Update, wg *sync.WaitGroup) {
	defer wg.Done()

	respChan := make(chan tgbotapi.Chattable)

	defer close(respChan)

	go RunHandler(upd, respChan)

	select {
	case resp := <-respChan:
		b.SendMessage(resp)

	case <-ctx.Done():
		return
	}

}
