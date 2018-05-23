package controllers

import (
	"gitlab.lisbontechhub.com/A-Team/challenger-chat/chaterrors"
	"gitlab.lisbontechhub.com/A-Team/challenger-chat/models"
)

/*
 Public Functions
*/

// CreateMessageController function that controls the flow to create a new message
// between two users.
func CreateMessageController(text string, sender string, receiver string) (*models.Message, error) {
	if len(text) == 0 {
		return nil, chaterrors.ErrTextIsEmpty{}
	}

	if len(sender) == 0 {
		return nil, chaterrors.ErrSenderNameIsEmpty{Username: sender}
	}

	if len(receiver) == 0 {
		return nil, chaterrors.ErrReceiverNameIsEmpty{Username: receiver}
	}

	if !CheckIfUserExists(sender) {
		return nil, chaterrors.ErrSenderDoesNotExist{Username: sender}
	}

	if !CheckIfUserExists(receiver) {
		return nil, chaterrors.ErrReceiverDoesNotExist{Username: receiver}
	}

	if !checkIfUsersAreFriends(sender, receiver) {
		return nil, chaterrors.ErrClientsAreNotFriends{Receiver: receiver, Sender: sender}
	}

	message, err := models.CreateMessage(text, sender, receiver)

	return message, err
}

// GetClientChatLogController function that controls the flow to get the chat log
// between two users.
func GetClientChatLogController(sender, receiver string) ([]models.Message, error) {
	if len(sender) == 0 {
		return nil, chaterrors.ErrSenderNameIsEmpty{Username: sender}
	}
	if len(receiver) == 0 {
		return nil, chaterrors.ErrReceiverNameIsEmpty{Username: receiver}
	}

	if !CheckIfUserExists(sender) {
		return nil, chaterrors.ErrSenderDoesNotExist{Username: sender}
	}

	if !CheckIfUserExists(receiver) {
		return nil, chaterrors.ErrReceiverDoesNotExist{Username: receiver}
	}

	messages := models.GetClientChatLog(sender, receiver)

	return messages, nil
}

/*
 Private Functions
*/
