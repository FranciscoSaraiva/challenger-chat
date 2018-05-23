package models

import (
	"fmt"
	"strings"

	"gitlab.lisbontechhub.com/A-Team/challenger-chat/chaterrors"
	"gitlab.lisbontechhub.com/A-Team/challenger-chat/config"

	"gopkg.in/mgo.v2/bson"
)

// Client of the chat
type Client struct {
	Name    string
	Status  string
	Friends []string
}

/*
 Public Functions
*/

// CreateClient function that creates a new Client.
// It creates a new Client if the name chosen isn't taken already.
// If it already exists a new Client isn't created.
// It functions as the register.
func CreateClient(name string) (Client, error) {
	collection := config.GetCollection("Client")

	client := new(Client)
	client.Name = name
	client.Status = "Online"
	client.Friends = nil

	err := collection.Insert(client)
	consoleOutputClient(*client, "create")
	return *client, err
}

// LoginClient function that loads the client. Functions as the login.
func LoginClient(name string) (Client, error) {
	collection := config.GetCollection("Client")
	query := bson.M{"name": name}

	client := new(Client)
	collection.Find(query).One(client)
	client.Status = "Online"

	err := collection.Update(query, client)
	chaterrors.CheckError(err)

	consoleOutputClient(*client, "login")

	return *client, err
}

// LogoutClient function that sets the client as logged out.
func LogoutClient(name string) (*Client, error) {
	collection := config.GetCollection("Client")
	query := bson.M{"name": name}

	client := new(Client)
	collection.Find(query).One(client)
	client.Status = "Offline"

	err := collection.Update(query, client)
	chaterrors.CheckError(err)

	return client, err
}

// GetClient function that returns an instance of a Client via its username.
func GetClient(name string) (*Client, error) {
	collection := config.GetCollection("Client")
	query := bson.M{"name": name}

	client := new(Client)
	collection.Find(query).One(client)

	return client, nil
}

// ChangeClientStatus function that changes the status of a client.
// The status are only between Online, Away and Busy.
func ChangeClientStatus(name, status string) (*Client, error) {
	collection := config.GetCollection("Client")
	query := bson.M{"name": name}

	client := new(Client)
	collection.Find(query).One(client)
	client.Status = status

	err := collection.Update(query, client)
	chaterrors.CheckError(err)

	return client, err
}

// ClientAddFriend function that adds a friend to a user.
// Both must exist to add each other.
func ClientAddFriend(user string, friend string) (*Client, error) {
	collection := config.GetCollection("Client")
	roomName := generateRoomName()
	fmt.Println("******************* ROOM NAME: " + roomName + " ****************")

	//Client side
	client := new(Client)
	query := bson.M{"name": user}
	collection.Find(query).One(client)
	client.Friends = append(client.Friends, friend+"-"+roomName)
	err := collection.Update(query, client)
	chaterrors.CheckError(err)

	//Friend side
	clientF := new(Client)
	query = bson.M{"name": friend}
	collection.Find(query).One(clientF)
	clientF.Friends = append(clientF.Friends, user+"-"+roomName)
	err = collection.Update(query, clientF)
	chaterrors.CheckError(err)

	return clientF, nil
}

// GetFriends function that returns an array of Clients with the information
// of a user's friends, such as username and current status.
func GetFriends(username string) ([]*Client, error) {
	var friendList []*Client
	collection := config.GetCollection("Client")
	query := bson.M{"name": username}

	client := new(Client)
	collection.Find(query).One(client)

	for _, v := range client.Friends {
		collection = config.GetCollection("Client")

		name := strings.Split(v, "-")

		query = bson.M{"name": name[0]}
		friend := new(Client)
		collection.Find(query).One(friend)

		friendList = append(friendList, friend)
		fmt.Println(name[0])
	}
	return friendList, nil
}

/**
Private Functions
*/

// generateRoomName function that generates a room name for two clients to
// talk inside the Socket.
func generateRoomName() string {
	roomName := bson.NewObjectId().Hex()
	return roomName
}

// consoleOutputClient function that returns information on the console, depending on the
// user's action. It shows information about a client and what action was made about them.
func consoleOutputClient(client Client, action string) {
	fmt.Println("-------------------------")

	switch action {
	case "create":
		action = "|Client Created|"
		break
	case "login":
		action = " |Client Login|"
		break
	}
	fmt.Println("    " + action + "")

	fmt.Println("*************************")
	fmt.Println("Name   -> " + client.Name)
	fmt.Println("Status -> " + client.Status)
	fmt.Println(">Friends<")
	for _, v := range client.Friends {
		fmt.Print(v + " ")
	}

	fmt.Println("\n*************************\n")
}
