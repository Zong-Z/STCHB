package texts

import (
	"encoding/json"
	"io/ioutil"
	"stchb/logger"
)

const File = "configs/texts.json"

type Text struct {
	Commands struct {
		Start struct {
			Command string `json:"command"`
			Text    string `json:"text"`
		} `json:"start"`
		Help struct {
			Command string `json:"command"`
			Text    string `json:"text"`
		} `json:"help"`
		Chatting struct {
			Start string `json:"start"`
			Stop  string `json:"stop"`
		} `json:"chatting"`
		Settings struct {
			Command string `json:"command"`
		} `json:"settings"`
		Me struct {
			Command string `json:"command"`
		} `json:"me"`
		BecomeAdministrator struct {
			Command string `json:"command"`
		} `json:"become_administrator"`
		Unknown struct {
			Text string `json:"text"`
		} `json:"unknown"`
	} `json:"commands"`
	ParseMode string `json:"parse_mode"`
	Chat      struct {
		NotRegistered             string `json:"not_registered"`
		InterlocutorSearchStarted string `json:"interlocutor_search_begun"`
		ChatFound                 string `json:"chat_found"`
		AlreadyInChat             string `json:"already_in_chat"`
		NotInChat                 string `json:"not_in_chat"`
		ChatEnded                 string `json:"chat_ended"`
	} `json:"chat"`
	InlineKeyboardMarkup struct {
		Opened  string `json:"opened"`
		Changed string `json:"changed"`
		Closed  string `json:"closed"`
	} `json:"inline_keyboard_markup"`
}

var text Text

func init() {
	b, err := ioutil.ReadFile(File)
	if err != nil {
		logger.ForWarning(err.Error())
	}

	if err := json.Unmarshal(b, &text); err != nil {
		logger.ForWarning(err.Error())
	}
}

func GetTexts() Text {
	return text
}
