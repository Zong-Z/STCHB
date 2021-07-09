package markups

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

const (
	AgeInlineKeyboardMarkupText     = "Age"
	AgeInlineKeyboardMarkupName     = "AGE"
	AgeInlineKeyboardMarkupCallback = AgeInlineKeyboardMarkupName
	AgeInlineKeyboardMarkupPrefix   = AgeInlineKeyboardMarkupName + "_"

	SixteenOrLessText     = "Sixteen or less"
	SixteenOrLessCallback = "SIXTEEN-OR-LESS"

	FromSixteenToEighteenText     = "From sixteen to eighteen"
	FromSixteenToEighteenCallback = "FROM-SIXTEEN-TO-EIGHTEEN"

	FromEighteenToTwentyOneText     = "From eighteen to twenty one"
	FromEighteenToTwentyOneCallback = "FROM-EIGHTEEN-TO-TWENTY-ONE"
)

var ownAgeInlineKeyboardMarkup = InlineKeyboardMarkup{
	Name: OwnPrefix + AgeInlineKeyboardMarkupName,
	InlineKeyboardMarkup: tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(SixteenOrLessText,
				SettingsInlineKeyboardMarkupPrefix+OwnPrefix+AgeInlineKeyboardMarkupPrefix+SixteenOrLessCallback),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(FromSixteenToEighteenText,
				SettingsInlineKeyboardMarkupPrefix+OwnPrefix+AgeInlineKeyboardMarkupPrefix+FromSixteenToEighteenCallback),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(FromEighteenToTwentyOneText,
				SettingsInlineKeyboardMarkupPrefix+OwnPrefix+AgeInlineKeyboardMarkupPrefix+FromEighteenToTwentyOneCallback),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(DoesNotMatterText,
				SettingsInlineKeyboardMarkupPrefix+OwnPrefix+AgeInlineKeyboardMarkupPrefix+DoesNotMatterCallback,
			),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(GoBackText,
				SettingsInlineKeyboardMarkupPrefix+OwnPrefix+AgeInlineKeyboardMarkupPrefix+GoBackCallback),
		),
	}},
}

var interlocutorAgeInlineKeyboardMarkup = InlineKeyboardMarkup{
	Name: InterlocutorPrefix + AgeInlineKeyboardMarkupName,
	InlineKeyboardMarkup: tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(SixteenOrLessText,
				SettingsInlineKeyboardMarkupPrefix+InterlocutorPrefix+AgeInlineKeyboardMarkupPrefix+SixteenOrLessCallback),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(FromSixteenToEighteenText,
				SettingsInlineKeyboardMarkupPrefix+InterlocutorPrefix+AgeInlineKeyboardMarkupPrefix+FromSixteenToEighteenCallback),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(FromEighteenToTwentyOneText,
				SettingsInlineKeyboardMarkupPrefix+InterlocutorPrefix+AgeInlineKeyboardMarkupPrefix+FromEighteenToTwentyOneCallback),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(DoesNotMatterText,
				SettingsInlineKeyboardMarkupPrefix+InterlocutorPrefix+AgeInlineKeyboardMarkupPrefix+DoesNotMatterCallback,
			),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(GoBackText,
				SettingsInlineKeyboardMarkupPrefix+InterlocutorPrefix+AgeInlineKeyboardMarkupPrefix+GoBackCallback),
		),
	}},
}
