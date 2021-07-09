package markups

import (
	"errors"
	"fmt"
	"stchb/database"
	"stchb/logger"
	"stchb/texts"
	"strings"

	"github.com/go-redis/redis/v8"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Settings struct{}

const (
	SettingsInlineKeyboardMarkupText     = "Settings"
	SettingsInlineKeyboardMarkupName     = "SETTINGS"
	SettingsInlineKeyboardMarkupCallback = SettingsInlineKeyboardMarkupName
	SettingsInlineKeyboardMarkupPrefix   = SettingsInlineKeyboardMarkupName + "_"

	DoesNotMatterText     = "Does not matter"
	DoesNotMatterName     = "DOES NOT MATTER"
	DoesNotMatterCallback = DoesNotMatterName

	GoBackText     = "<-Back"
	GoBackName     = "BACK"
	GoBackCallback = GoBackName
)

const (
	OwnPrefixText = "Own"
	OwnPrefix     = "OWN_"

	InterlocutorPrefixText = "Interlocutor"
	InterlocutorPrefix     = "INTERLOCUTOR_"
)

var SettingsInlineKeyboardMarkup = InlineKeyboardMarkup{
	Name: SettingsInlineKeyboardMarkupName,
	InlineKeyboardMarkup: tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(OwnPrefixText+" "+AgeInlineKeyboardMarkupText,
				SettingsInlineKeyboardMarkupPrefix+OwnPrefix+AgeInlineKeyboardMarkupCallback),
			tgbotapi.NewInlineKeyboardButtonData(OwnPrefixText+" "+AgeInlineKeyboardMarkupText,
				SettingsInlineKeyboardMarkupPrefix+InterlocutorPrefix+AgeInlineKeyboardMarkupCallback),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(OwnPrefixText+" "+CityInlineKeyboardMarkupText,
				SettingsInlineKeyboardMarkupPrefix+OwnPrefix+CityInlineKeyboardMarkupCallback),
			tgbotapi.NewInlineKeyboardButtonData(InterlocutorPrefixText+" "+CityInlineKeyboardMarkupText,
				SettingsInlineKeyboardMarkupPrefix+InterlocutorPrefix+CityInlineKeyboardMarkupCallback),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(OwnPrefixText+" "+SexInlineKeyboardMarkupText,
				SettingsInlineKeyboardMarkupPrefix+OwnPrefix+SexInlineKeyboardMarkupCallback),
			tgbotapi.NewInlineKeyboardButtonData(InterlocutorPrefixText+" "+SexInlineKeyboardMarkupText,
				SettingsInlineKeyboardMarkupPrefix+InterlocutorPrefix+SexInlineKeyboardMarkupCallback),
		),
	}},
}

var SettingsInlineKeyboardMarkups = InlineKeyboardMarkups{
	SettingsInlineKeyboardMarkup,
	ownAgeInlineKeyboardMarkup, interlocutorAgeInlineKeyboardMarkup,
	ownCityInlineKeyboardMarkup, interlocutorCityInlineKeyboardMarkup,
	ownSexInlineKeyboardMarkup, interlocutorSexInlineKeyboardMarkup,
}

func (Settings) IsMarkupRequestByCallbackQuery(callbackQuery tgbotapi.CallbackQuery) bool {
	return SettingsInlineKeyboardMarkups.FindInlineKeyboardMarkup(callbackQuery.Data) != nil
}

func (Settings) SendMarkupByCallbackQueryRequest(callbackQuery tgbotapi.CallbackQuery, botAPI *tgbotapi.BotAPI) error {
	_, err := database.DB.GetUser(callbackQuery.From.ID)
	if err != nil && err.Error() == redis.Nil.Error() {
		msg := tgbotapi.MessageConfig{
			BaseChat:  tgbotapi.BaseChat{ChatID: int64(callbackQuery.From.ID)},
			Text:      texts.GetTexts().Chat.NotRegistered,
			ParseMode: texts.GetTexts().ParseMode,
		}

		if _, err := botAPI.Send(msg); err != nil {
			return err
		}

		logger.ForInfo(fmt.Sprintf("User %d, not registered.", callbackQuery.From.ID))

		return nil
	} else if err != nil {
		return err
	}

	inlineKeyboardMarkup := SettingsInlineKeyboardMarkups.FindInlineKeyboardMarkup(callbackQuery.Data)
	if inlineKeyboardMarkup == nil {
		return errors.New("inlineKeyboardMarkup was not found, unknown callback")
	}

	for _, buttons := range inlineKeyboardMarkup.InlineKeyboard {
		for _, button := range buttons {
			if strings.Contains(*button.CallbackData, AgeInlineKeyboardMarkupPrefix) ||
				strings.Contains(*button.CallbackData, CityInlineKeyboardMarkupPrefix) ||
				strings.Contains(*button.CallbackData, SexInlineKeyboardMarkupPrefix) {
				button.Text = fmt.Sprintf("->%s<-", button.Text)
			}
		}
	}

	if _, err := botAPI.Send(tgbotapi.NewEditMessageReplyMarkup(int64(callbackQuery.From.ID),
		callbackQuery.Message.MessageID, inlineKeyboardMarkup.InlineKeyboardMarkup)); err != nil {
		return err
	}

	_, err = botAPI.AnswerCallbackQuery(tgbotapi.NewCallback(callbackQuery.ID,
		texts.GetTexts().InlineKeyboardMarkup.Opened))
	return err
}

func (Settings) IsCallbackQueryRequestForChangeUserSettings(callbackQuery tgbotapi.CallbackQuery) bool {
	return strings.Contains(callbackQuery.Data, AgeInlineKeyboardMarkupPrefix) ||
		strings.Contains(callbackQuery.Data, CityInlineKeyboardMarkupPrefix) ||
		strings.Contains(callbackQuery.Data, SettingsInlineKeyboardMarkupPrefix)
}

func (Settings) ChangeUserSettingsByCallbackQueryRequest(callbackQuery tgbotapi.CallbackQuery, botAPI *tgbotapi.BotAPI) error {
	settings := Settings{}
	if err := func() error {
		user, err := database.DB.GetUser(callbackQuery.From.ID)
		if err != nil && err.Error() == redis.Nil.Error() {
			msg := tgbotapi.MessageConfig{
				BaseChat:  tgbotapi.BaseChat{ChatID: int64(callbackQuery.From.ID)},
				Text:      texts.GetTexts().Chat.NotRegistered,
				ParseMode: texts.GetTexts().ParseMode,
			}

			if _, err := botAPI.Send(msg); err != nil {
				return err
			}

			logger.ForInfo(fmt.Sprintf("User %d, not registered.", callbackQuery.From.ID))

			return errors.New("user not registered")
		} else if err != nil {
			return err
		}

		isRequestToChangeOwnSettings := settings.isRequestToChangeOwnSettings(callbackQuery.Data)
		dataToChange := strings.Replace(callbackQuery.Data, SettingsInlineKeyboardMarkupPrefix, "", -1)
		if isRequestToChangeOwnSettings {
			dataToChange = strings.Replace(dataToChange, OwnPrefix, "", -1)
		} else {
			dataToChange = strings.Replace(dataToChange, InterlocutorPrefix, "", -1)
		}

		switch {
		case settings.isRequestToChangeAge(callbackQuery.Data):
			dataToChange = strings.Replace(dataToChange, AgeInlineKeyboardMarkupPrefix, "", -1)
			if isRequestToChangeOwnSettings {
				user.Age = dataToChange
			} else {
				user.AgeOfTheInterlocutor = dataToChange
			}
		case settings.isRequestToChangeCity(callbackQuery.Data):
			dataToChange = strings.Replace(dataToChange, CityInlineKeyboardMarkupPrefix, "", -1)
			if isRequestToChangeOwnSettings {
				user.City = dataToChange
			} else {
				user.CityOfTheInterlocutor = dataToChange
			}
		case settings.isRequestToChangeSex(callbackQuery.Data):
			dataToChange = strings.Replace(dataToChange, SettingsInlineKeyboardMarkupPrefix, "", -1)
			if isRequestToChangeOwnSettings {
				user.Sex = dataToChange
			} else {
				user.SexOfTheInterlocutor = dataToChange
			}
		}

		if err := database.DB.SaveUser(*user); err != nil {
			return err
		}

		return nil
	}(); err != nil {
		return err
	}

	if _, err := botAPI.AnswerCallbackQuery(tgbotapi.NewCallback(callbackQuery.ID,
		texts.GetTexts().InlineKeyboardMarkup.Changed)); err != nil {
		return err
	}

	if _, err := botAPI.Send(tgbotapi.NewEditMessageReplyMarkup(int64(callbackQuery.From.ID),
		callbackQuery.Message.MessageID, SettingsInlineKeyboardMarkup.InlineKeyboardMarkup)); err != nil {
		return err
	}

	return nil
}

func (Settings) isRequestToChangeOwnSettings(callbackQueryData string) bool {
	return strings.Contains(callbackQueryData, OwnPrefix)
}

func (Settings) isRequestToChangeAge(callbackQueryData string) bool {

	return strings.Contains(callbackQueryData, AgeInlineKeyboardMarkupPrefix)
}

func (Settings) isRequestToChangeCity(callbackQueryData string) bool {
	return strings.Contains(callbackQueryData, CityInlineKeyboardMarkupPrefix)
}

func (Settings) isRequestToChangeSex(callbackQueryData string) bool {
	return strings.Contains(callbackQueryData, SettingsInlineKeyboardMarkupPrefix)
}
