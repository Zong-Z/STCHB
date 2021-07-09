package types

import "errors"

func NewChats(usersCount, buffer int) (*Chats, error) {
	if usersCount < 2 {
		return nil, errors.New("can't be less than two users in the chat")
	} else if buffer < 1 {
		return nil, errors.New("the buffer must be greater than 1")
	}

	chats := &Chats{
		Chats:      make([]Chat, 0),
		UsersCount: usersCount,
		Queue:      make(chan User, buffer),
	}

	go chats.StartAddingUsers()

	return chats, nil
}

func (c *Chats) StartAddingUsers() {
	for u := range c.Queue {
		if len(c.Chats) == 0 {
			c.AddChat(*NewChat())
			c.Chats[len(c.Chats)-1].AddUserToChat(u)
			continue
		}

		chats := c.FindSuitableChatsForUser(u)
		if chats != nil {
			chats[0].AddUserToChat(u)
			continue
		}

		c.AddChat(*NewChat())
		c.Chats[len(c.Chats)-1].AddUserToChat(u)
	}
}

func (c *Chats) AddChat(chat Chat) {
	c.Chats = append(c.Chats, chat)
}

func (c *Chats) FindSuitableChatsForUser(user User) []*Chat {
	chats := make([]*Chat, 0)
	for i := 0; i < len(c.Chats); i++ {
		usersCount := len(c.Chats[i].Users)
		if usersCount == c.UsersCount {
			continue
		}

		func(suitable ...func(User) bool) {
			u := c.Chats[i].Users[len(c.Chats[i].Users)-1]
			for _, f := range suitable {
				if !f(u) {
					return
				}
			}

			chats = append(chats, &c.Chats[i])
		}(user.IsSuitableAge, user.IsSuitableCity, user.IsSuitableSex)
	}

	if len(chats) != 0 {
		return chats
	}

	return nil
}

func (c *Chats) AddUserToQueue(user User) {
	c.Queue <- user
}

func (c *Chats) DeleteUserFromChat(userID int) {
	for i := 0; i < len(c.Chats); i++ {
		for j := 0; j < len(c.Chats[i].Users); j++ {
			if c.Chats[i].Users[j].ID == userID {
				c.Chats[i].Users[len(c.Chats[i].Users)-1],
					c.Chats[i].Users[i] = c.Chats[i].Users[i], c.Chats[i].Users[len(c.Chats[i].Users)-1]
				c.Chats[i].Users = c.Chats[i].Users[:len(c.Chats[i].Users)-1]
				return
			}
		}
	}
}

func (c *Chats) DeleteChatWithUser(userID int) {
	for i := 0; i < len(c.Chats); i++ {
		for j := 0; j < len(c.Chats[i].Users); j++ {
			if c.Chats[i].Users[j].ID == userID {
				c.Chats[len(c.Chats)-1], c.Chats[i] = c.Chats[i], c.Chats[len(c.Chats)-1]
				c.Chats = c.Chats[:len(c.Chats)-1]
				return
			}
		}
	}
}

func (c *Chats) GetInterlocutorsByUserID(userID int) []User {
	if !c.IsUserInChat(userID) {
		return nil
	}

	for i := 0; i < len(c.Chats); i++ {
		if c.Chats[i].IsUserInChat(userID) {
			return c.Chats[i].GetInterlocutorsByUserID(userID)
		}
	}

	return nil
}

func (c *Chats) IsUserInChat(userID int) bool {
	for i := 0; i < len(c.Chats); i++ {
		if c.Chats[i].IsUserInChat(userID) {
			return true
		}
	}
	return false
}
