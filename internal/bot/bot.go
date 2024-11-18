package bot

import (
	"context"
	"sync"

	"go_tgbot/internal/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	API    *tgbotapi.BotAPI
	Config *config.Config
}

func Init(config *config.Config) *Bot {
	return &Bot{
		Config: config,
	}
}

func (b *Bot) Run(ctx context.Context, wg *sync.WaitGroup) error {
	err := b.InitBotAPI()
	if err != nil {
		return err
	}

	updates := b.API.GetUpdatesChan(tgbotapi.UpdateConfig{})
	for update := range updates {
		wg.Add(1)
		go b.HandleUpdate(ctx, update, wg)
	}

	return nil
}

func (b *Bot) InitBotAPI() error {
	botAPI, err := tgbotapi.NewBotAPI(b.Config.APIToken)
	if err != nil {
		return err
	}

	b.API = botAPI

	return nil
}

func (b *Bot) SendMessage(msg tgbotapi.Chattable) {
	b.API.Request(msg)
}
