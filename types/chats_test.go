package types

import "testing"

func TestNewChats(t *testing.T) {
	chats, _ := NewChats(2, 10)
	if chats == nil {
		t.Errorf("Chat creation error: %v.", chats)
		t.Fail()
	}

	_, err := NewChats(0, 10)
	if err == nil {
		t.Errorf("Chat creation error: %v.", chats)
		t.Fail()
	}

	_, err = NewChats(2, 0)
	if err == nil {
		t.Errorf("Chat creation error: %v.", chats)
		t.Fail()
	}

	_, err = NewChats(0, 0)
	if err == nil {
		t.Errorf("Chat creation error: %v.", chats)
		t.Fail()
	}
}
