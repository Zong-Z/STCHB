package bot

import (
	"fmt"
	"stchb/bot/markups"
	"stchb/config"
	"stchb/logger"
	"stchb/texts"
	"stchb/types"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Updates struct{}

func (Updates) HandleUpdate(update tgbotapi.Update, chats *types.Chats, botAPI *tgbotapi.BotAPI) {
	updates := Updates{}
	switch {
	case update.CallbackQuery != nil:
		updates.handleCallbackQuery(*update.CallbackQuery, botAPI)
	case update.Message != nil:
		updates.handleMessage(*update.Message, chats, botAPI)
	}
}

func (Updates) handleCallbackQuery(callbackQuery tgbotapi.CallbackQuery, botAPI *tgbotapi.BotAPI) {
	settings := markups.Settings{}
	switch {
	case markups.IsCloseCallbackQuery(callbackQuery):
		if _, err := botAPI.Send(
			tgbotapi.NewDeleteMessage(int64(callbackQuery.From.ID), callbackQuery.Message.MessageID)); err != nil {
			logger.ForError(err.Error())
		}
	case settings.IsMarkupRequestByCallbackQuery(callbackQuery):
		if err := settings.SendMarkupByCallbackQueryRequest(callbackQuery, botAPI); err != nil {
			logger.ForError(err.Error())
		}
	case settings.IsCallbackQueryRequestForChangeUserSettings(callbackQuery):
		if err := settings.ChangeUserSettingsByCallbackQueryRequest(callbackQuery, botAPI); err != nil {
			logger.ForError(err.Error())
		}
	}
}

func (Updates) handleMessage(message tgbotapi.Message, chats *types.Chats, botAPI *tgbotapi.BotAPI) {
	updates := Updates{}
	command := message.IsCommand()
	switch {
	case command:
		updates.handleCommand(message, chats, botAPI)
	case !command && chats.IsUserInChat(message.From.ID): // Message to anther user.
		updates.handleMessageToAnotherUser(message, chats, botAPI)
	}
}

func (Updates) handleCommand(message tgbotapi.Message, chats *types.Chats, botAPI *tgbotapi.BotAPI) {
	commands := Commands{}
	switch message.Command() {
	case texts.GetTexts().Commands.Start.Command:
		if err := commands.Start(*types.NewUser(tgbotapi.User{
			ID:           message.From.ID,
			FirstName:    message.From.FirstName,
			LastName:     message.From.LastName,
			UserName:     message.From.UserName,
			LanguageCode: message.From.LanguageCode,
			IsBot:        message.From.IsBot}), botAPI); err != nil {
			logger.ForError(err.Error())
		}
	case texts.GetTexts().Commands.Help.Command:
		if err := commands.Help(message.From.ID, botAPI); err != nil {
			logger.ForError(err.Error())
		}
	case texts.GetTexts().Commands.Chatting.Start:
		if err := commands.StartChatting(message.From.ID, chats, botAPI); err != nil {
			logger.ForError(err.Error())
		}
	case texts.GetTexts().Commands.Chatting.Stop:
		if err := commands.StartChatting(message.From.ID, chats, botAPI); err != nil {
			logger.ForError(err.Error())
		}
	case texts.GetTexts().Commands.Settings.Command:
		if err := commands.Settings(message.From.ID, botAPI); err != nil {
			logger.ForError(err.Error())
		}
	case texts.GetTexts().Commands.Me.Command:
		if err := commands.Me(message.From.ID, botAPI); err != nil {
			logger.ForError(err.Error())
		}
	default:
		if _, err := botAPI.Send(tgbotapi.MessageConfig{
			BaseChat:  tgbotapi.BaseChat{ChatID: int64(message.From.ID)},
			Text:      texts.GetTexts().Commands.Unknown.Text,
			ParseMode: texts.GetTexts().ParseMode,
		}); err != nil {
			logger.ForError(err.Error())
		}
	}
}

func (Updates) handleMessageToAnotherUser(message tgbotapi.Message, chats *types.Chats, botAPI *tgbotapi.BotAPI) {
	var messageToAnotherUser tgbotapi.Chattable
	interlocutors := chats.GetInterlocutorsByUserID(message.From.ID)
	if interlocutors == nil {
		if _, err := botAPI.Send(tgbotapi.MessageConfig{
			BaseChat:  tgbotapi.BaseChat{ChatID: int64(message.From.ID)},
			Text:      texts.GetTexts().Chat.NotInChat,
			ParseMode: texts.GetTexts().ParseMode,
		}); err != nil {
			logger.ForError(err.Error())
		}
	}

	for i, interlocutor := range interlocutors {
		switch {
		case message.Audio != nil:
			messageToAnotherUser = tgbotapi.NewAudioShare(int64(interlocutor.ID), message.Audio.FileID)
		case message.Document != nil:
			messageToAnotherUser = tgbotapi.NewDocumentShare(int64(interlocutor.ID), message.Document.FileID)
		case message.Animation != nil:
			messageToAnotherUser = tgbotapi.NewAnimationShare(int64(interlocutor.ID), message.Animation.FileID)
		case message.Photo != nil:
			messageToAnotherUser = tgbotapi.NewPhotoShare(int64(interlocutor.ID), (*message.Photo)[0].FileID)
		case message.Sticker != nil:
			messageToAnotherUser = tgbotapi.NewStickerShare(int64(interlocutor.ID), message.Sticker.FileID)
		case message.Video != nil:
			messageToAnotherUser = tgbotapi.NewVideoShare(int64(interlocutor.ID), message.Video.FileID)
		case message.VideoNote != nil:
			messageToAnotherUser = tgbotapi.NewVideoNoteShare(int64(interlocutor.ID), message.VideoNote.Length,
				message.VideoNote.FileID)
		case message.Voice != nil:
			messageToAnotherUser = tgbotapi.NewVoiceShare(int64(interlocutor.ID), message.Voice.FileID)
		default:
			if config.GetConfig().Chat.Users > 2 { // If there are more than two users in the chat.
				messageToAnotherUser = tgbotapi.MessageConfig{
					BaseChat:  tgbotapi.BaseChat{ChatID: int64(interlocutor.ID)},
					Text:      fmt.Sprintf("*INTERLOCUTOR %d:* %s", i+1, message.Text),
					ParseMode: "MARKDOWN",
				}

				break
			}

			messageToAnotherUser = tgbotapi.MessageConfig{
				BaseChat: tgbotapi.BaseChat{ChatID: int64(interlocutor.ID)},
				Text:     message.Text,
			}
		}

		if _, err := botAPI.Send(messageToAnotherUser); err != nil {
			logger.ForError(err.Error())
		}
	}
}

func (Updates) SetPolling(offset, limit, timeout int, botAPI *tgbotapi.BotAPI) (tgbotapi.UpdatesChannel, error) {
	updates, err := botAPI.GetUpdatesChan(tgbotapi.UpdateConfig{
		Offset:  offset,
		Limit:   limit,
		Timeout: timeout,
	})
	if err != nil {
		return nil, err
	}

	return updates, nil
}

func (Updates) SetWebhook(webhook string, botAPI *tgbotapi.BotAPI) (tgbotapi.UpdatesChannel, error) {
	_, err := botAPI.SetWebhook(tgbotapi.NewWebhook(webhook))
	if err != nil {
		return nil, err
	}

	return botAPI.ListenForWebhook("/"), nil
}
