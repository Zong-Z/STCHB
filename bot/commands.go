package bot

import (
	"crypto/sha256"
	"fmt"
	"stchb/bot/markups"
	"stchb/config"
	"stchb/database"
	"stchb/texts"
	"stchb/types"

	"github.com/go-redis/redis/v8"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Commands struct{}

func (Commands) Start(user types.User, botAPI *tgbotapi.BotAPI) error {
	_, err := database.DB.GetUser(user.ID)
	if err != nil && err.Error() == redis.Nil.Error() {
		if err := database.DB.SaveUser(user); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	_, err = botAPI.Send(tgbotapi.MessageConfig{
		BaseChat:  tgbotapi.BaseChat{ChatID: int64(user.ID)},
		Text:      texts.GetTexts().Commands.Start.Text,
		ParseMode: texts.GetTexts().ParseMode,
	})
	return err
}

func (Commands) Help(chatID int, botAPI *tgbotapi.BotAPI) error {
	_, err := botAPI.Send(tgbotapi.MessageConfig{
		BaseChat:  tgbotapi.BaseChat{ChatID: int64(chatID)},
		Text:      texts.GetTexts().Commands.Help.Text,
		ParseMode: texts.GetTexts().ParseMode,
	})
	return err
}

func (Commands) StartChatting(userID int, chats *types.Chats, botAPI *tgbotapi.BotAPI) error {
	msg := tgbotapi.MessageConfig{
		BaseChat:  tgbotapi.BaseChat{ChatID: int64(userID)},
		Text:      texts.GetTexts().Chat.AlreadyInChat,
		ParseMode: texts.GetTexts().ParseMode,
	}

	if !chats.IsUserInChat(userID) {
		user, err := database.DB.GetUser(userID)
		if err != nil && err.Error() == redis.Nil.Error() {
			msg.Text = texts.GetTexts().Chat.NotRegistered
			if _, err := botAPI.Send(msg); err != nil {
				return err
			}

			return nil
		} else if err != nil {
			return err
		}

		chats.AddUserToQueue(*user)

		msg.Text = texts.GetTexts().Chat.InterlocutorSearchStarted
	}

	if _, err := botAPI.Send(msg); err != nil {
		return err
	}

	interlocutors := chats.GetInterlocutorsByUserID(userID)
	if interlocutors != nil && len(interlocutors)+1 >= 2 {
		msg := tgbotapi.MessageConfig{
			Text:      texts.GetTexts().Chat.ChatFound,
			ParseMode: texts.GetTexts().ParseMode,
		}

		for i := 0; i < len(interlocutors); i++ {
			msg.ChatID = int64(interlocutors[i].ID)
			if _, err := botAPI.Send(msg); err != nil {
				return err
			}
		}

		msg.ChatID = int64(userID)
		if _, err := botAPI.Send(msg); err != nil {
			return err
		}
	}

	return nil
}

func (Commands) StopChatting(userID int, chats *types.Chats, botAPI *tgbotapi.BotAPI) error {
	msg := tgbotapi.MessageConfig{
		ParseMode: texts.GetTexts().ParseMode,
		Text:      texts.GetTexts().Chat.ChatEnded,
	}

	interlocutors := chats.GetInterlocutorsByUserID(userID)
	if len(interlocutors) == 1 {
		chats.DeleteChatWithUser(userID)
		msg.ChatID = int64(interlocutors[0].ID)
	} else if interlocutors == nil && chats.IsUserInChat(userID) {
		chats.DeleteChatWithUser(userID)
		msg.ChatID = int64(userID)
	} else {
		chats.DeleteUserFromChat(userID)
		msg.ParseMode = "MARKDOWN"
		for _, interlocutor := range interlocutors {
			msg.ChatID = int64(interlocutor.ID)
			msg.Text = fmt.Sprintf("*INTERLOCUTOR %d LEFT THE CHAT*", func() int {
				var number int
				for i, interlocutor := range chats.GetInterlocutorsByUserID(interlocutor.ID) {
					if userID == interlocutor.ID {
						number = i + 1
						break
					}
				}

				return number
			}())

			if _, err := botAPI.Send(msg); err != nil {
				return err
			}
		}

		return nil
	}

	_, err := botAPI.Send(msg)
	return err
}

func (Commands) Settings(userID int, botAPI *tgbotapi.BotAPI) error {
	_, err := botAPI.Send(tgbotapi.MessageConfig{
		BaseChat:  tgbotapi.BaseChat{ChatID: int64(userID), ReplyMarkup: markups.SettingsInlineKeyboardMarkup},
		Text:      markups.SexInlineKeyboardMarkupText,
		ParseMode: texts.GetTexts().ParseMode,
	})
	return err
}

func (Commands) Me(userID int, botAPI *tgbotapi.BotAPI) error {
	u, err := database.DB.GetUser(userID)
	if err != nil && err.Error() == redis.Nil.Error() {
		if _, err := botAPI.Send(tgbotapi.MessageConfig{
			BaseChat:  tgbotapi.BaseChat{ChatID: int64(userID)},
			Text:      texts.GetTexts().Chat.NotRegistered,
			ParseMode: texts.GetTexts().ParseMode,
		}); err != nil {
			return err
		}

		return nil
	} else if err != nil {
		return err
	}

	_, err = botAPI.Send(tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{ChatID: int64(userID)},
		Text: fmt.Sprintf(
			"*Information about you.*\n"+
				"*Your age:* %s.\n"+
				"*Interlocutor age:* %s.\n"+
				"*Your city:* %s.\n"+
				"*Interlocutor city:* %s.\n"+
				"*Your sex:* %s.\n"+
				"*Sex of the interlocutor:* %s.",
			u.Age, u.AgeOfTheInterlocutor,
			u.City, u.CityOfTheInterlocutor,
			u.Sex, u.SexOfTheInterlocutor),
		ParseMode: "MARKDOWN",
	})
	return err
}

func (Commands) BecomeAdministrator(userID int, password []byte, botAPI *tgbotapi.BotAPI) error {
	if config.GetConfig().Bot.AdministratorPasswordSHA256 != fmt.Sprintf("%x", sha256.Sum256(password)) {
		return InvalidPassword
	}

	u, err := database.DB.GetUser(userID)
	if err != nil && err.Error() == redis.Nil.Error() {
		if _, err := botAPI.Send(tgbotapi.MessageConfig{
			BaseChat:  tgbotapi.BaseChat{ChatID: int64(userID)},
			Text:      texts.GetTexts().Chat.NotRegistered,
			ParseMode: texts.GetTexts().ParseMode,
		}); err != nil {
			return err
		}

		return nil
	} else if err != nil {
		return err
	}

	u.IsAdmin = true
	return database.DB.SaveUser(*u)
}
