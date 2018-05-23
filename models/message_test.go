package models

import (
	"testing"
)

func TestCreateMessageCorrect(t *testing.T) {
	text := "Oi"
	sender := "Francisco"
	receiver := "Luis"
	message, err := CreateMessage(text, sender, receiver)
	if checkError(err) {
		t.Errorf("Expected no error and got [%v]", err)
	}
	if message.Text != "Oi" {
		t.Errorf("Expected [%s] got [%s]", text, message.Text)
	}
	if message.Sender != "Francisco" {
		t.Errorf("Expected [%s] got [%s]", sender, message.Sender)
	}
	if message.Receiver != "Luis" {
		t.Errorf("Expected [%s] got [%s]", receiver, message.Receiver)
	}
}

func TestGetClientChatLogIsNotEmpty(t *testing.T) {
	sender := "Luis"
	receiver := "Francisco"
	messages := GetClientChatLog(sender, receiver)

	if len(messages) == 0 {
		t.Errorf("Expected higher than zero but got [%d]", len(messages))
	}
}

func TestGetClientChatLogIsEmpty(t *testing.T) {
	sender := "John Doe"
	receiver := "James Rolfe"
	messages := GetClientChatLog(sender, receiver)

	if len(messages) != 0 {
		t.Errorf("Expected 0 got [%d]", len(messages))
	}
}
