package markups

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type InlineKeyboardMarkup struct {
	Name                          string `json:"name"`
	tgbotapi.InlineKeyboardMarkup `json:"tgbotapi_._inline_keyboard_markup"`
}

type InlineKeyboardMarkups []InlineKeyboardMarkup

func (markups InlineKeyboardMarkups) FindInlineKeyboardMarkup(inlineKeyboardMarkupName string) *InlineKeyboardMarkup {
	for _, markup := range markups {
		if markup.Name != inlineKeyboardMarkupName {
			continue
		}

		var inlineKeyboardMarkup InlineKeyboardMarkup
		inlineKeyboardMarkup.InlineKeyboard = make([][]tgbotapi.InlineKeyboardButton, len(markup.InlineKeyboard))
		for i, buttons := range markup.InlineKeyboard {
			for _, button := range buttons {
				inlineKeyboardMarkup.InlineKeyboard[i] = append(inlineKeyboardMarkup.InlineKeyboard[i],
					tgbotapi.InlineKeyboardButton{
						Text: button.Text, URL: button.URL, CallbackData: button.CallbackData, Pay: button.Pay,
						SwitchInlineQuery: button.SwitchInlineQuery, CallbackGame: button.CallbackGame,
						SwitchInlineQueryCurrentChat: button.SwitchInlineQueryCurrentChat,
					},
				)
			}
		}

		return &inlineKeyboardMarkup
	}

	return nil
}

const InlineKeyboardMarkupCloseCallback = "CLOSE"

func IsCloseCallbackQuery(callbackQuery tgbotapi.CallbackQuery) bool {
	return strings.Contains(callbackQuery.Data, InlineKeyboardMarkupCloseCallback)
}
