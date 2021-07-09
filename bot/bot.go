package bot

type BotError string

func (e BotError) Error() string { return string(e) }

const InvalidPassword = BotError("invalid password")
