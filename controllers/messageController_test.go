package controllers

import (
	"log"
	"reflect"
	"testing"

	"gitlab.lisbontechhub.com/A-Team/challenger-chat/chaterrors"
	"gitlab.lisbontechhub.com/A-Team/challenger-chat/config"
	"gitlab.lisbontechhub.com/A-Team/challenger-chat/models"
)

/**
TestCreateMessageController
*/

func TestCreateMessageControllerCorrect(t *testing.T) {
	dropCollections()
	/** FOR IT TO WORK **/
	clientS, _ := models.CreateClient("John Doe")
	clientR, _ := models.CreateClient("James Rolfe")
	models.ClientAddFriend(clientS.Name, clientR.Name)
	/**/
	sender := clientS.Name
	receiver := clientR.Name
	text := "Oi"
	message, err := CreateMessageController(text, sender, receiver)

	if checkError(err) {
		t.Errorf("Expected no error but got Error: [%v]", err)
	}
	if message.Text != text {
		t.Errorf("Expected the text message to be [%s] but it was [%s]", text, message.Text)
	}
	if message.Sender != sender {
		t.Errorf("Expected the Sender to be [%s] but it was [%s]", sender, message.Sender)
	}
	if message.Receiver != receiver {
		t.Errorf("Expected the Receiver to be [%s] but it was [%s]", receiver, message.Receiver)
	}
}

func TestCreateMessageControllerErrTextIsEmpty(t *testing.T) {
	dropCollections()
	/** FOR IT TO WORK **/
	clientS, _ := models.CreateClient("John Doe")
	clientR, _ := models.CreateClient("James Rolfe")
	models.ClientAddFriend(clientS.Name, clientR.Name)
	/**/
	sender := clientS.Name
	receiver := clientR.Name
	text := ""
	message, err := CreateMessageController(text, sender, receiver)

	if message != nil {
		t.Errorf("Expected the Message to be nil but it was [%v]", message)
	}
	if !chaterrors.IsErrTextIsEmpty(err) {
		t.Errorf("Expected error to be of type 'ErrTextIsEmpty' but got Error: [%v]", reflect.TypeOf(err))
	}
}

func TestCreateMessageControllerErrSenderNameIsEmpty(t *testing.T) {
	dropCollections()
	/** FOR IT TO WORK **/
	clientS, _ := models.CreateClient("John Doe")
	clientR, _ := models.CreateClient("James Rolfe")
	models.ClientAddFriend(clientS.Name, clientR.Name)
	/**/
	sender := ""
	receiver := clientR.Name
	text := "Oi"
	message, err := CreateMessageController(text, sender, receiver)

	if message != nil {
		t.Errorf("Expected the Message to be nil but it was [%v]", message)
	}
	if !chaterrors.IsErrSenderNameIsEmpty(err) {
		t.Errorf("Expected error to be of type 'ErrSenderNameIsEmpty' but got Error: [%v]", reflect.TypeOf(err))
	}
}

func TestCreateMessageControllerErrReceiverNameIsEmpty(t *testing.T) {
	dropCollections()
	/** FOR IT TO WORK **/
	clientS, _ := models.CreateClient("John Doe")
	clientR, _ := models.CreateClient("James Rolfe")
	models.ClientAddFriend(clientS.Name, clientR.Name)
	/**/
	sender := clientS.Name
	receiver := ""
	text := "Oi"
	message, err := CreateMessageController(text, sender, receiver)

	if message != nil {
		t.Errorf("Expected the Message to be nil but it was [%v]", message)
	}
	if !chaterrors.IsErrReceiverNameIsEmpty(err) {
		t.Errorf("Expected error to be of type 'ErrReceiverNameIsEmpty' but got Error: [%v]", reflect.TypeOf(err))
	}
}

func TestCreateMessageControllerErrSenderDoesNotExist(t *testing.T) {
	dropCollections()
	/** FOR IT TO WORK **/
	clientR, _ := models.CreateClient("James Rolfe")
	models.ClientAddFriend("John Doe", clientR.Name)
	/**/
	sender := "John Doe"
	receiver := clientR.Name
	text := "Oi"
	message, err := CreateMessageController(text, sender, receiver)

	if message != nil {
		t.Errorf("Expected the Message to be nil but it was [%v]", message)
	}
	if !chaterrors.IsErrSenderDoesNotExist(err) {
		t.Errorf("Expected error to be of type 'ErrSenderDoesNotExist' but got Error: [%v]", reflect.TypeOf(err))
	}
}

func TestCreateMessageControllerErrReceiverDoesNotExist(t *testing.T) {
	dropCollections()
	/** FOR IT TO WORK **/
	clientS, _ := models.CreateClient("John Doe")
	models.ClientAddFriend(clientS.Name, "James Rolfe")
	/**/
	sender := clientS.Name
	receiver := "James Rolfe"
	text := "Oi"

	message, err := CreateMessageController(text, sender, receiver)

	if message != nil {
		t.Errorf("Expected the Message to be nil but it was [%v]", message)
	}
	if !chaterrors.IsErrReceiverDoesNotExist(err) {
		t.Errorf("Expected error to be of type 'ErrReceiverDoesNotExist' but got Error: [%v]", reflect.TypeOf(err))
	}
}

func TestCreateMessageControllerErrClientAreNotFriends(t *testing.T) {
	dropCollections()
	/** FOR IT TO WORK **/
	clientS, _ := models.CreateClient("John Doe")
	clientR, _ := models.CreateClient("James Rolfe")
	/**/
	sender := clientS.Name
	receiver := clientR.Name
	text := "Oi"
	message, err := CreateMessageController(text, sender, receiver)

	if message != nil {
		t.Errorf("Expected the Message to be nil but it was [%v]", message)
	}
	if !chaterrors.IsErrClientsAreNotFriends(err) {
		t.Errorf("Expected error to be of type 'ErrClientsAreNotFriends' but got Error: [%v]", reflect.TypeOf(err))
	}
}

/**
TestGetClientChatLogController
*/

func TestGetClientChatLogControllerCorrect(t *testing.T) {
	dropCollections()
	/** FOR IT TO WORK **/
	clientS, _ := models.CreateClient("John Doe")
	clientR, _ := models.CreateClient("James Rolfe")
	models.ClientAddFriend(clientS.Name, clientR.Name)
	receiver := "James Rolfe"
	firstText := "Oi-1"
	secondText := "Oi-2"
	thirdText := "Oi-3"
	models.CreateMessage(firstText, clientS.Name, clientR.Name)
	models.CreateMessage(secondText, clientS.Name, clientR.Name)
	models.CreateMessage(thirdText, clientS.Name, clientR.Name)
	/**/

	messages, err := GetClientChatLogController(clientS.Name, receiver)

	// ERROR
	if checkError(err) {
		t.Errorf("Expected no error but got Error: [%v]", err)
	}
	//SENDER
	if messages[0].Sender != clientS.Name {
		t.Errorf("Expected the Sender of the first text to be [%s] but it was [%s]", clientS.Name, messages[0].Sender)
	}
	if messages[1].Sender != clientS.Name {
		t.Errorf("Expected the Sender of the second text to be [%s] but it was [%s]", clientS.Name, messages[1].Sender)
	}
	if messages[2].Sender != clientS.Name {
		t.Errorf("Expected the Sender of the third text to be [%s] but it was [%s]", clientS.Name, messages[2].Sender)
	}
	//RECEIVER
	if messages[0].Receiver != receiver {
		t.Errorf("Expected the Receiver of the first text to be [%s] but it was [%s]", clientR.Name, messages[0].Receiver)
	}
	if messages[1].Receiver != receiver {
		t.Errorf("Expected the Receiver of the second text to be [%s] but it was [%s]", clientR.Name, messages[1].Receiver)
	}
	if messages[2].Receiver != receiver {
		t.Errorf("Expected the Receiver of the third text to be [%s] but it was [%s]", clientR.Name, messages[2].Receiver)
	}
	//SIZE
	if len(messages) != 3 {
		t.Errorf("Expected the lenght of the messages to be 3 but it was [%d]", len(messages))
	}
	//TEXT
	if messages[0].Text != firstText {
		t.Errorf("Expected the first text to be [%s] but it was [%s]", firstText, messages[0].Text)
	}
	if messages[1].Text != secondText {
		t.Errorf("Expected the second text to be [%s] but it was [%s]", secondText, messages[1].Text)
	}
	if messages[2].Text != thirdText {
		t.Errorf("Expected the third text to be [%s] but it was [%s]", thirdText, messages[2].Text)
	}

}

func TestGetClientChatLogControllerErrSenderNameIsEmpty(t *testing.T) {
	dropCollections()
	/** FOR IT TO WORK **/
	clientR, _ := models.CreateClient("James Rolfe")
	models.ClientAddFriend("John Doe", clientR.Name)
	sender := ""
	receiver := "James Rolfe"
	firstText := "Oi-1"
	secondText := "Oi-2"
	thirdText := "Oi-3"
	models.CreateMessage(firstText, sender, clientR.Name)
	models.CreateMessage(secondText, sender, clientR.Name)
	models.CreateMessage(thirdText, sender, clientR.Name)
	/**/

	messages, err := GetClientChatLogController(sender, receiver)

	if messages != nil {
		t.Errorf("Expected the Messages to be nil but it was [%v]", messages)
	}
	if !chaterrors.IsErrSenderNameIsEmpty(err) {
		t.Errorf("Expected error to be of type 'ErrSenderNameIsEmpty' but got Error: [%v]", reflect.TypeOf(err))
	}
}

func TestGetClientChatLogControllerErrReceiverNameIsEmpty(t *testing.T) {
	dropCollections()
	/** FOR IT TO WORK **/
	clientS, _ := models.CreateClient("John Doe")
	clientR, _ := models.CreateClient("James Rolfe")
	models.ClientAddFriend(clientS.Name, clientR.Name)
	sender := "John Doe"
	receiver := ""
	firstText := "Oi-1"
	secondText := "Oi-2"
	thirdText := "Oi-3"
	models.CreateMessage(firstText, sender, clientR.Name)
	models.CreateMessage(secondText, sender, clientR.Name)
	models.CreateMessage(thirdText, sender, clientR.Name)
	/**/

	messages, err := GetClientChatLogController(sender, receiver)

	if messages != nil {
		t.Errorf("Expected the Messages to be nil but it was [%v]", messages)
	}
	if !chaterrors.IsErrReceiverNameIsEmpty(err) {
		t.Errorf("Expected error to be of type 'ErrReceiverNameIsEmpty' but got Error: [%v]", reflect.TypeOf(err))
	}
}

func TestGetClientChatLogControllerErrSenderDoesNotExist(t *testing.T) {
	dropCollections()
	/** FOR IT TO WORK **/
	clientR, _ := models.CreateClient("James Rolfe")
	models.ClientAddFriend("John Doe", clientR.Name)
	sender := "John Doe"
	receiver := "James Rolfe"
	firstText := "Oi-1"
	secondText := "Oi-2"
	thirdText := "Oi-3"
	models.CreateMessage(firstText, sender, clientR.Name)
	models.CreateMessage(secondText, sender, clientR.Name)
	models.CreateMessage(thirdText, sender, clientR.Name)
	/**/

	messages, err := GetClientChatLogController(sender, receiver)

	if messages != nil {
		t.Errorf("Expected the Message to be nil but it was [%v]", messages)
	}
	if !chaterrors.IsErrSenderDoesNotExist(err) {
		t.Errorf("Expected error to be of type 'ErrSenderDoesNotExist' but got Error: [%v]", reflect.TypeOf(err))
	}
}

func TestGetClientChatLogControllerErrReceiverDoesNotExist(t *testing.T) {
	dropCollections()
	/** FOR IT TO WORK **/
	clientS, _ := models.CreateClient("John Doe")
	models.ClientAddFriend(clientS.Name, "James Rolfe")
	sender := "John Doe"
	receiver := "James Rolfe"
	firstText := "Oi-1"
	secondText := "Oi-2"
	thirdText := "Oi-3"
	models.CreateMessage(firstText, sender, receiver)
	models.CreateMessage(secondText, sender, receiver)
	models.CreateMessage(thirdText, sender, receiver)
	/**/

	messages, err := GetClientChatLogController(clientS.Name, receiver)

	if messages != nil {
		t.Errorf("Expected the Messages to be nil but it was [%v]", messages)
	}
	if !chaterrors.IsErrReceiverDoesNotExist(err) {
		t.Errorf("Expected error to be of type 'ErrReceiverDoesNotExist' but got Error: [%v]", reflect.TypeOf(err))
	}
}

//CheckError function that checks for an error and logs what error occurred.
func checkError(error error) bool {
	if error != nil {
		log.Printf("%v", error)
		return true
	}
	return false
}

//Drop collections for testing purposes
func dropCollections() {
	config.GetCollection("Message").DropCollection()
	config.GetCollection("Client").DropCollection()
	config.GetCollection("Session").DropCollection()
}
