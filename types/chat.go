package types

import "github.com/google/uuid"

func NewChat() *Chat {
	return &Chat{
		Users: make([]User, 0),
		ID:    uuid.New().String(),
	}
}

func (c *Chat) AddUserToChat(user User) {
	c.Users = append(c.Users, user)
}

func (c *Chat) DeleteUserFromChat(userID int) {
	for i := 0; i < len(c.Users); i++ {
		if c.Users[i].ID == userID {
			c.Users[len(c.Users)-1], c.Users[i] = c.Users[i], c.Users[len(c.Users)-1]
			c.Users = c.Users[:len(c.Users)-1]
			break
		}
	}
}

func (c *Chat) GetInterlocutorsByUserID(userID int) []User {
	if !c.IsUserInChat(userID) {
		return nil
	}

	users := make([]User, 0)
	for i := 0; i < len(c.Users); i++ {
		if c.Users[i].ID != userID {
			users = append(users, c.Users[i])
		}
	}

	if len(users) != 0 {
		return users
	}

	return nil
}

func (c *Chat) IsUserInChat(userID int) bool {
	for i := 0; i < len(c.Users); i++ {
		if c.Users[i].ID == userID {
			return true
		}
	}

	return false
}
