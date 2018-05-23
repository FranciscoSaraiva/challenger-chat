package models

import (
	"fmt"
	"time"

	"gitlab.lisbontechhub.com/A-Team/challenger-chat/config"

	"gopkg.in/mgo.v2/bson"
)

// Message clients send each other
type Message struct {
	Sender    string
	Receiver  string
	Text      string
	Timestamp int64
}

/*
 Public Functions
*/

// CreateMessage function that receives the information needed to create a new message
// between two users. Both users must exist in the database to be created.
func CreateMessage(text string, sender string, receiver string) (*Message, error) {
	message := new(Message)
	message.Text = text
	message.Sender = sender
	message.Receiver = receiver
	message.Timestamp = config.GetTimeNowTimeStamp()

	collection := config.GetCollection("Message")
	err := collection.Insert(message)
	consoleOutputMessage(*message, "send")

	return message, err
}

// GetClientChatLog function that receives a username and a friend name and returns the conversations
// between them.
func GetClientChatLog(sender, receiver string) []Message {
	collection := config.GetCollection("Message")

	var allMessages []Message

	queryClient := bson.M{"sender": sender, "receiver": receiver}
	queryFriend := bson.M{"sender": receiver, "receiver": sender}
	collection.Find(bson.M{"$or": []bson.M{queryClient, queryFriend}}).Sort("timestamp").All(&allMessages)

	return allMessages
}

/*
 Private Functions
*/

// consoleOutputClient function that returns information on the console, depending on the
// user's action. It shows information about a client and what action was made about them.
func consoleOutputMessage(message Message, action string) {
	fmt.Println("-------------------------")

	switch action {
	case "send":
		action = "|New Message|"
		break
	}
	fmt.Println("\n    " + action + "")

	fmt.Println("*************************")
	fmt.Println("Sender   -> " + message.Sender)
	fmt.Println("Receiver -> " + message.Receiver)
	fmt.Println("Time     -> " + time.Unix(message.Timestamp, 0).String())
	fmt.Println(">Message<")
	fmt.Println(message.Text + "\n")

	fmt.Println("*************************\n")
}
