package types

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

const UserNil = "USER: NIL"

type User struct {
	tgbotapi.User         `json:"tgbotapi_._user"`
	IsAdmin               bool   `json:"is_admin"`
	Age                   string `json:"age"`
	City                  string `json:"city"`
	Sex                   string `json:"sex"`
	AgeOfTheInterlocutor  string `json:"age_of_the_interlocutor"`
	CityOfTheInterlocutor string `json:"city_of_the_interlocutor"`
	SexOfTheInterlocutor  string `json:"sex_of_the_interlocutor"`
}

type Chat struct {
	Users []User `json:"users"`
	ID    string `json:"id"`
}

type Queue chan User

type Chats struct {
	Chats      []Chat `json:"chats"`
	UsersCount int    `json:"users_count"` // Maximum number of users in one chat.
	Queue      `json:"queue"`
}
