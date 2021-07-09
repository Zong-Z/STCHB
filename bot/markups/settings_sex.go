package markups

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

const (
	SexInlineKeyboardMarkupText     = "Sex"
	SexInlineKeyboardMarkupName     = "SEX"
	SexInlineKeyboardMarkupCallback = SexInlineKeyboardMarkupName
	SexInlineKeyboardMarkupPrefix   = SexInlineKeyboardMarkupName + "_"

	MaleText       = "Male"
	MaleCallback   = "MALE"
	FemaleText     = "Female"
	FemaleCallback = "FEMALE"
)

var ownSexInlineKeyboardMarkup = InlineKeyboardMarkup{
	Name: OwnPrefix + SexInlineKeyboardMarkupName,
	InlineKeyboardMarkup: tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(MaleText,
				SettingsInlineKeyboardMarkupPrefix+OwnPrefix+MaleCallback),
			tgbotapi.NewInlineKeyboardButtonData(FemaleText,
				SettingsInlineKeyboardMarkupPrefix+OwnPrefix+FemaleCallback),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(DoesNotMatterText,
				SettingsInlineKeyboardMarkupPrefix+OwnPrefix+SexInlineKeyboardMarkupPrefix+DoesNotMatterCallback,
			),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(GoBackText,
				SettingsInlineKeyboardMarkupPrefix+OwnPrefix+SexInlineKeyboardMarkupPrefix+GoBackCallback),
		),
	}},
}

var interlocutorSexInlineKeyboardMarkup = InlineKeyboardMarkup{
	Name: InterlocutorPrefix + SexInlineKeyboardMarkupName,
	InlineKeyboardMarkup: tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(MaleText,
				SettingsInlineKeyboardMarkupPrefix+InterlocutorPrefix+MaleCallback),
			tgbotapi.NewInlineKeyboardButtonData(FemaleText,
				SettingsInlineKeyboardMarkupPrefix+InterlocutorPrefix+FemaleCallback),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(DoesNotMatterText,
				SettingsInlineKeyboardMarkupPrefix+InterlocutorPrefix+SexInlineKeyboardMarkupPrefix+DoesNotMatterCallback,
			),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(GoBackText,
				SettingsInlineKeyboardMarkupPrefix+InterlocutorPrefix+SexInlineKeyboardMarkupPrefix+GoBackCallback),
		),
	}},
}
