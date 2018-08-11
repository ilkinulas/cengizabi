package main

import (
	"flag"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ilkinulas/cengizabi/command"
	"github.com/ilkinulas/cengizabi/config"
	"log"
	"os"
)

var (
	configFile string
)

func main() {
	parseFlags()
	logger := log.New(os.Stdout, "", log.LstdFlags)
	cfg, err := config.Load(configFile)
	if err != nil {
		logger.Fatalf("Failed to load config file	. %v", err)
	}
	logger.Printf("Starting Cengiz Abi %v", GetHumanVersion())
	bot, err := tgbotapi.NewBotAPI(cfg.BotApiToken)
	if err != nil {
		logger.Fatalf("Failed to start bot. %v", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	log.Printf("Telegram bot %s is ready. Listening for updates...", bot.Self.UserName)

	cmdRegistry := command.NewRegistry()

	for update := range updates {
		msg, err := handleUpdate(update, cmdRegistry, logger)
		if err != nil {
			logger.Printf("Failed to handle update. %v", err)
			continue
		}
		bot.Send(msg)
	}
}

func handleUpdate(update tgbotapi.Update, cmdRegistry *command.Registry, logger *log.Logger) (*tgbotapi.MessageConfig, error) {
	if update.Message == nil {
		return nil, fmt.Errorf("update.message is nil")
	}
	text := update.Message.Text
	cmdOut, err := cmdRegistry.HandleMessage(text)
	if err != nil {
		return nil, fmt.Errorf("command exec failed. %v", err)
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, cmdOut.Text)
	//msg.ReplyToMessageID = update.Message.MessageID
	logger.Printf("Replying with %v", cmdOut.Text)
	return &msg, nil
}

func parseFlags() {
	flag.StringVar(&configFile, "config", "config.toml", "path to configuration file.")
	flag.Parse()
}
