package main

import (
	"net/http"
	"stchb/bot"
	"stchb/config"
	"stchb/logger"
	"stchb/types"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	go func() {
		logger.ForError(http.ListenAndServe(":"+config.GetConfig().Bot.Port, nil))
	}()

	logger.ForInfo("Authorized on account.")
	botAPI, err := tgbotapi.NewBotAPI(config.GetConfig().Bot.Token)
	if err != nil {
		logger.ForError(err.Error())
	}

	logger.ForInfo("Bot have created successfully.")

	updates := bot.Updates{}
	var updatesChannel tgbotapi.UpdatesChannel
	if !strings.EqualFold(config.GetConfig().Bot.Webhook, "") {
		updatesChannel, err = updates.SetWebhook(config.GetConfig().Bot.Webhook, botAPI)
		if err != nil {
			logger.ForError(err.Error())
		}
	} else {
		updatesChannel, err = updates.SetPolling(config.GetConfig().Bot.Polling.Offset,
			config.GetConfig().Bot.Polling.Limit, config.GetConfig().Bot.Polling.Timeout, botAPI)
		if err != nil {
			logger.ForError(err.Error())
		}
	}

	chats, err := types.NewChats(config.GetConfig().Chat.Queue, config.GetConfig().Chat.Users)
	if err != nil {
		logger.ForError(err.Error())
	}

	for update := range updatesChannel {
		updates.HandleUpdate(update, chats, botAPI)
	}
}
