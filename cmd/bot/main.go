package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"go_tgbot/internal/bot"
	"go_tgbot/internal/config"
)

func main() {
	configPath := flag.String("c", "./config.yaml", "path to bot config")
	flag.Parse()

	cfg := &config.Config{}
	config.GetConfiguration(*configPath, cfg)

	myBot := bot.Init(cfg)

	ctx, cancel := context.WithCancel(context.Background())

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	var wg sync.WaitGroup

	go myBot.Run(ctx, &wg)

	<-signalChan
	fmt.Println("\nShutting down server...")

	cancel()

	wg.Wait()

}
