package controllers

import (
	"fmt"
	"strings"

	"gitlab.lisbontechhub.com/A-Team/challenger-chat/chaterrors"
	"gitlab.lisbontechhub.com/A-Team/challenger-chat/config"
	"gitlab.lisbontechhub.com/A-Team/challenger-chat/models"

	"gopkg.in/mgo.v2/bson"
)

/*
 Public Functions
*/

// StartClientController function that starts the Client entry process.
// If the Client's name already exists, then it proceeds to loading his information.
// If the Client's name doesn't exist, then it creates a new Client and loads his information.
func StartClientController(name string) (*models.Client, error) {
	if len(name) == 0 {
		return nil, chaterrors.ErrClientLoginNameIsEmpty{Username: name}
	}

	if len(name) < 3 {
		return nil, chaterrors.ErrClientLoginNameIsShort{Username: name}
	}

	for _, v := range name {
		if v >= '0' && v <= '9' {
			return nil, chaterrors.ErrClientLoginNameHasNumeric{Username: name}
		}
	}

	if CheckIfUserExists(name) {
		client, err := models.LoginClient(name)
		return &client, err
	}

	client, err := models.CreateClient(name)
	return &client, err
}

// LogoutClientController function that controls the flow to logout the client from the application.
func LogoutClientController(name string) (*models.Client, error) {
	if len(name) == 0 {
		return nil, chaterrors.ErrClientLoginNameIsEmpty{Username: name}
	}

	if !CheckIfUserExists(name) {
		return nil, chaterrors.ErrClientNotFound{Username: name}
	}
	client, err := models.LogoutClient(name)
	return client, err
}

// ClientAddFriendController function that controls the flow for a client to add a friend.
func ClientAddFriendController(user string, friend string) error {
	if len(user) == 0 {
		return chaterrors.ErrClientNameEmpty{Username: user}
	}
	if len(friend) == 0 {
		return chaterrors.ErrEmptyFriendName{Friendname: friend}
	}

	if !CheckIfUserExists(friend) {
		return chaterrors.ErrFriendNotFound{Friendname: friend}
	}

	if checkIfUsersAreFriends(user, friend) {
		return chaterrors.ErrUsersAreFriends{Username: user, Friendname: friend}
	}
	_, err := models.ClientAddFriend(user, friend)

	return err
}

// GetFriendsController function that controls the flow to getting the friends of a client.
func GetFriendsController(username string) ([]*models.Client, error) {
	if len(username) == 0 {
		return nil, chaterrors.ErrClientNameEmpty{Username: username}
	}

	if !CheckIfUserExists(username) {
		return nil, chaterrors.ErrClientNotFound{Username: username}
	}

	friends, err := models.GetFriends(username)
	return friends, err
}

// GetClientController function that controls the flow of getting the information of a client.
func GetClientController(name string) (*models.Client, error) {
	if len(name) == 0 {
		return nil, chaterrors.ErrClientNameEmpty{Username: name}
	}

	if !CheckIfUserExists(name) {
		return nil, chaterrors.ErrClientNotFound{Username: name}
	}
	client, err := models.GetClient(name)
	return client, err
}

// ChangeClientStatusController function that controls the flow to change a client's status.
func ChangeClientStatusController(name, status string) (*models.Client, error) {
	status = strings.ToLower(status)
	status = strings.Title(status)

	if len(name) == 0 {
		return nil, chaterrors.ErrClientNameEmpty{Username: name}
	}

	if !CheckIfUserExists(name) {
		return nil, chaterrors.ErrClientNotFound{Username: name}
	}

	fmt.Println("CURRENT STATUS: " + status)
	fmt.Println(status == "Online")
	fmt.Println(status == "Busy")
	fmt.Println(status == "Away")

	if status != "Online" && status != "Busy" && status != "Away" {
		return nil, chaterrors.ErrInvalidStatus{Status: status}
	}
	client, err := models.ChangeClientStatus(name, status)
	return client, err
}

// CheckIfUserExists function that checks if a user with the name already exists in the database.
// If it exists, a user was already created and returns true.
// If it doesn't exist, a user still wasn't created and returns false.
func CheckIfUserExists(name string) bool {
	users := config.GetCollection("Client")
	count, err := users.Find(bson.M{"name": name}).Count()
	if !chaterrors.CheckError(err) && count == 1 {
		return true
	}

	return false
}

/*
 Private Functions
*/

// checkIfUsersAreFriends function that checks if two users are already friends.
func checkIfUsersAreFriends(user, friend string) bool {
	var userList []string

	collection := config.GetCollection("Client")

	//Client side
	client := new(models.Client)
	query := bson.M{"name": user}
	collection.Find(query).One(client)
	userList = client.Friends

	for _, v := range userList {
		friendName := strings.Split(v, "-")
		if friendName[0] == friend {
			return true
		}
	}
	return false
}
