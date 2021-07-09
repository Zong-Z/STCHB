package markups

import (
	"stchb/cities"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	CityInlineKeyboardMarkupText     = "City"
	CityInlineKeyboardMarkupName     = "CITY"
	CityInlineKeyboardMarkupCallback = CityInlineKeyboardMarkupName
	CityInlineKeyboardMarkupPrefix   = CityInlineKeyboardMarkupName + "_"
)

var ownCityInlineKeyboardMarkup, interlocutorCityInlineKeyboardMarkup InlineKeyboardMarkup

func init() {
	ownCityInlineKeyboardMarkup.Name = OwnPrefix + CityInlineKeyboardMarkupName
	ownCityInlineKeyboardMarkup.InlineKeyboard = make([][]tgbotapi.InlineKeyboardButton, 0)
	for _, city := range cities.Cities {
		ownCityInlineKeyboardMarkup.InlineKeyboard = append(ownCityInlineKeyboardMarkup.InlineKeyboard,
			tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(city,
				SettingsInlineKeyboardMarkupPrefix+OwnPrefix+CityInlineKeyboardMarkupPrefix+strings.ToUpper(city))),
		)
	}

	interlocutorCityInlineKeyboardMarkup.Name = InterlocutorPrefix + CityInlineKeyboardMarkupName
	interlocutorCityInlineKeyboardMarkup.InlineKeyboard = make([][]tgbotapi.InlineKeyboardButton, 0)
	for _, city := range cities.Cities {
		interlocutorCityInlineKeyboardMarkup.InlineKeyboard = append(ownCityInlineKeyboardMarkup.InlineKeyboard,
			tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(city,
				SettingsInlineKeyboardMarkupPrefix+InterlocutorPrefix+CityInlineKeyboardMarkupPrefix+strings.ToUpper(city))),
		)
	}

	ownCityInlineKeyboardMarkup.InlineKeyboard = append(ownCityInlineKeyboardMarkup.InlineKeyboard,
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(
			DoesNotMatterText,
			SettingsInlineKeyboardMarkupPrefix+OwnPrefix+SexInlineKeyboardMarkupPrefix+DoesNotMatterCallback)),
	)
	ownCityInlineKeyboardMarkup.InlineKeyboard = append(ownCityInlineKeyboardMarkup.InlineKeyboard,
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(GoBackText,
			SettingsInlineKeyboardMarkupPrefix+OwnPrefix+SexInlineKeyboardMarkupPrefix+GoBackCallback)),
	)

	interlocutorCityInlineKeyboardMarkup.InlineKeyboard = append(ownCityInlineKeyboardMarkup.InlineKeyboard,
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(
			DoesNotMatterText,
			SettingsInlineKeyboardMarkupPrefix+OwnPrefix+SexInlineKeyboardMarkupPrefix+DoesNotMatterCallback)),
	)
	interlocutorCityInlineKeyboardMarkup.InlineKeyboard = append(ownCityInlineKeyboardMarkup.InlineKeyboard,
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(GoBackText,
			SettingsInlineKeyboardMarkupPrefix+OwnPrefix+SexInlineKeyboardMarkupPrefix+GoBackCallback)),
	)
}
