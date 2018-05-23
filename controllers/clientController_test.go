package controllers

import (
	"testing"

	"gitlab.lisbontechhub.com/A-Team/challenger-chat/chaterrors"
	"gitlab.lisbontechhub.com/A-Team/challenger-chat/models"
)

func TestStartClientControllerCreate(t *testing.T) {
	dropCollections()
	/** TO WORK **/
	username := "James Rolfe"
	/**/

	client, err := StartClientController(username)

	if checkError(err) {
		t.Errorf("Expected no error but got Error: [%v]", err)
	}
	if client.Name != username {
		t.Errorf("Expected the name to be [%s] but it was [%s]", username, client.Name)
	}
	if len(client.Friends) != 0 {
		t.Errorf("Expected the friend list length to be 0 but it was [%d]", len(client.Friends))
	}
	if client.Status != "Online" {
		t.Errorf("Expected the status of the client to be 'Online' but it was [%s]", client.Status)
	}
}

func TestStartClientControllerLogin(t *testing.T) {
	dropCollections()
	/** TO WORK **/
	username := "James Rolfe"
	models.CreateClient(username)
	models.ChangeClientStatus(username, "Offline")
	/**/

	client, err := StartClientController(username)

	if checkError(err) {
		t.Errorf("Expected no error but got Error: [%v]", err)
	}
	if client.Name != username {
		t.Errorf("Expected the name to be [%s] but it was [%s]", username, client.Name)
	}
	if len(client.Friends) != 0 {
		t.Errorf("Expected the friend list length to be 0 but it was [%d]", len(client.Friends))
	}
	if client.Status != "Online" {
		t.Errorf("Expected the status of the client to be 'Online' but it was [%s]", client.Status)
	}
}

func TestStartClientControllerErrClientLoginNameIsEmpty(t *testing.T) {
	dropCollections()
	/** TO WORK **/
	username := ""
	/**/

	client, err := StartClientController(username)

	if client != nil {
		t.Errorf("Expected the client to be nil but it was created")
	}
	if !chaterrors.IsErrClientLoginNameIsEmpty(err) {
		t.Errorf("Expected Error: 'ErrClientLoginNameIsEmpty' but got error [%v]", err)
	}
}

func TestStartClientControllerErrClientLoginNameIsShort(t *testing.T) {
	dropCollections()
	/** TO WORK **/
	username := "ok"
	/**/

	client, err := StartClientController(username)

	if client != nil {
		t.Errorf("Expected the client to be nil but it was created")
	}
	if !chaterrors.IsErrClientLoginNameIsShort(err) {
		t.Errorf("Expected Error: 'ErrClientLoginNameIsShort' but got error [%v]", err)
	}
}

func TestStartClientControllerErrClientLoginNameHasNumeric(t *testing.T) {
	dropCollections()
	/** TO WORK **/
	username := "User123"
	/**/

	client, err := StartClientController(username)

	if client != nil {
		t.Errorf("Expected the client to be nil but it was created")
	}
	if !chaterrors.IsErrClientLoginNameHasNumeric(err) {
		t.Errorf("Expected Error: 'ErrClientLoginNameHasNumeric' but it was [%v]", err)
	}
}

func TestLogoutClientControllerCorrect(t *testing.T) {
	dropCollections()
	/** TO WORK **/
	username := "James Rolfe"
	StartClientController(username)
	/**/

	client, err := LogoutClientController(username)

	if checkError(err) {
		t.Errorf("Expected no error but got Error: [%v]", err)
	}
	if client.Name != username {
		t.Errorf("Expected the name to be [%s] but it was [%s]", username, client.Name)
	}
	if len(client.Friends) != 0 {
		t.Errorf("Expected the friend list length to be 0 but it was [%d]", len(client.Friends))
	}
	if client.Status != "Offline" {
		t.Errorf("Expected the status of the client to be 'Offline' but it was [%s]", client.Status)
	}
}

func TestLogoutClientControllerErrClientLoginNameIsEmpty(t *testing.T) {
	dropCollections()
	/** TO WORK **/
	username := ""
	/**/

	client, err := LogoutClientController(username)

	if client != nil {
		t.Errorf("Expected the client to be nil but it was created")
	}
	if !chaterrors.IsErrClientLoginNameIsEmpty(err) {
		t.Errorf("Expected Error: 'ErrClientLoginNameIsEmpty' but it was [%v]", err)
	}
}

func TestLogoutClientControllerErrClientNotFound(t *testing.T) {
	dropCollections()
	/** TO WORK **/
	username := "James Rolfe"
	/**/

	client, err := LogoutClientController(username)

	if client != nil {
		t.Errorf("Expected the client to be nil but it was created")
	}
	if !chaterrors.IsErrClientNotFound(err) {
		t.Errorf("Expected Error: 'ErrClientNotFound' but it was [%v]", err)
	}
}

func TestClientAddFriendControllerCorrect(t *testing.T) {
	dropCollections()
	/** TO WORK **/
	user := "James Rolfe"
	friend := "John Doe"
	models.CreateClient(user)
	models.CreateClient(friend)
	/**/

	err := ClientAddFriendController(user, friend)

	if checkError(err) {
		t.Errorf("Expected no error but got Error: [%v]", err)
	}

	friends, _ := models.GetFriends(user)

	if len(friends) != 1 {
		t.Errorf("Expected the length of the friends list to be 1 but it was [%d]", len(friends))
	}
	if friends[0].Name != friend {
		t.Errorf("Expected the friend name to be [%s] but it was [%s]", friend, friends[0].Name)
	}
}

func TestClientAddFriendControllerErrClientNameEmpty(t *testing.T) {
	dropCollections()
	/** TO WORK **/
	user := ""
	friend := "John Doe"
	models.CreateClient(user)
	models.CreateClient(friend)
	/**/

	err := ClientAddFriendController(user, friend)

	if !chaterrors.IsErrClientNameEmpty(err) {
		t.Errorf("Expected Error: 'ErrClientNameEmpty' but it was [%v]", err)
	}
}

func TestClientAddFriendControllerErrEmptyFriendName(t *testing.T) {
	dropCollections()
	/** TO WORK **/
	user := "James Rolfe"
	friend := ""
	models.CreateClient(user)
	models.CreateClient(friend)
	/**/

	err := ClientAddFriendController(user, friend)

	if !chaterrors.IsErrEmptyFriendName(err) {
		t.Errorf("Expected Error: 'ErrEmptyFriendName' but got [%v]", err)
	}
}

func TestClientAddFriendControllerErrFriendNotFound(t *testing.T) {
	dropCollections()
	/** TO WORK **/
	user := "James Rolfe"
	friend := "John Doe"
	models.CreateClient(user)
	/**/

	err := ClientAddFriendController(user, friend)

	if !chaterrors.IsErrFriendNotFound(err) {
		t.Errorf("Expected Error: 'ErrFriendNotFound' but got [%v]", err)
	}
}

func TestClientAddFriendControllerErrUsersAreFriends(t *testing.T) {
	dropCollections()
	/** TO WORK **/
	user := "James Rolfe"
	friend := "John Doe"
	models.CreateClient(user)
	models.CreateClient(friend)
	models.ClientAddFriend(user, friend)
	/**/

	err := ClientAddFriendController(user, friend)

	if !chaterrors.IsErrUsersAreFriends(err) {
		t.Errorf("Expected the Error: 'ErrUsersAreFriends' but got [%v]", err)
	}
}

func TestGetFriendsControllerCorrect(t *testing.T) {
	dropCollections()
	/** TO WORK **/
	user := "James Rolfe"
	friend := "John Doe"
	models.CreateClient(user)
	models.CreateClient(friend)
	models.ClientAddFriend(user, friend)
	/**/

	friends, err := GetFriendsController(user)

	if checkError(err) {
		t.Errorf("Expected no error but got error [%v]", err)
	}
	if len(friends) != 1 {
		t.Errorf("Expected length of friends to be 1 but it was [%d]", len(friends))
	}
	if friends[0].Name != friend {
		t.Errorf("Expected the name of the friend to be [%s] but it was [%s]", friend, friends[0].Name)
	}
}

func TestGetFriendsControllerErrClientNameEmpty(t *testing.T) {
	dropCollections()
	/** TO WORK **/
	user := ""
	friend := "John Doe"
	models.CreateClient(user)
	models.CreateClient(friend)
	models.ClientAddFriend(user, friend)
	/**/

	friends, err := GetFriendsController(user)

	if friends != nil {
		t.Errorf("Expected friends list to be nil but it was created")
	}
	if !chaterrors.IsErrClientNameEmpty(err) {
		t.Errorf("Expected Error:'ErrClientNameEmpty' but got [%v]", err)
	}
}

func TestGetFriendsControllerErrClientNotFound(t *testing.T) {
	dropCollections()
	/** TO WORK **/
	user := "James Rolfe"
	friend := "John Doe"
	models.CreateClient(friend)
	models.ClientAddFriend(user, friend)
	/**/

	friends, err := GetFriendsController(user)

	if friends != nil {
		t.Errorf("Expected friends list to be nil but it was created")
	}
	if !chaterrors.IsErrClientNotFound(err) {
		t.Errorf("Expected Error:'ErrClientNotFound' but got [%v]", err)
	}
}

func TestGetClientControllerCorrect(t *testing.T) {
	dropCollections()
	/** TO WORK **/
	username := "James Rolfe"
	models.CreateClient(username)
	/**/

	client, err := GetClientController(username)

	if checkError(err) {
		t.Errorf("Expected no error but got Error: [%v]", err)
	}
	if client.Name != username {
		t.Errorf("Expected the name to be [%s] but it was [%s]", username, client.Name)
	}
	if len(client.Friends) != 0 {
		t.Errorf("Expected the friend list length to be 0 but it was [%d]", len(client.Friends))
	}
	if client.Status != "Online" {
		t.Errorf("Expected the status of the client to be 'Online' but it was [%s]", client.Status)
	}
}

func TestGetClientControllerErrClientNameEmpty(t *testing.T) {
	dropCollections()
	/** TO WORK **/
	username := ""
	models.CreateClient(username)
	/**/

	client, err := GetClientController(username)

	if client != nil {
		t.Errorf("Expected the client to be nil but it was created")
	}
	if !chaterrors.IsErrClientNameEmpty(err) {
		t.Errorf("Expected Error: 'ErrClientNameEmpty' [%v]", err)
	}
}

func TestGetClientControllerErrClientNotFound(t *testing.T) {
	dropCollections()
	/** TO WORK **/
	username := "James Rolfe"
	/**/

	client, err := GetClientController(username)

	if client != nil {
		t.Errorf("Expected the client to be nil but it was created")
	}
	if !chaterrors.IsErrClientNotFound(err) {
		t.Errorf("Expected Error: 'ErrClientNotFound' [%v]", err)
	}
}

func TestChangeClientStatusControllerCorrect(t *testing.T) {
	dropCollections()
	/** TO WORK **/
	username := "James Rolfe"
	models.CreateClient(username)
	/**/

	client, err := ChangeClientStatusController(username, "Busy")

	if checkError(err) {
		t.Errorf("Expected no error but got Error: [%v]", err)
	}
	if client.Name != username {
		t.Errorf("Expected the name to be [%s] but it was [%s]", username, client.Name)
	}
	if len(client.Friends) != 0 {
		t.Errorf("Expected the friend list length to be 0 but it was [%d]", len(client.Friends))
	}
	if client.Status != "Busy" {
		t.Errorf("Expected the status of the client to change to 'Busy' but it was [%s]", client.Status)
	}
}

func TestChangeClientStatusControllerErrClientNameEmpty(t *testing.T) {
	dropCollections()
	/** TO WORK **/
	username := ""
	models.CreateClient(username)
	/**/

	client, err := ChangeClientStatusController(username, "Busy")

	if client != nil {
		t.Errorf("Expected client to be nil but it was created")
	}
	if !chaterrors.IsErrClientNameEmpty(err) {
		t.Errorf("Expected Error: 'ErrClientNameEmpty' but got [%v]", err)
	}
}

func TestChangeClientStatusControllerErrClientNotFound(t *testing.T) {
	dropCollections()
	/** TO WORK **/
	username := "James Rolfe"
	/**/

	client, err := ChangeClientStatusController(username, "Busy")

	if client != nil {
		t.Errorf("Expected client to be nil but it was created")
	}
	if !chaterrors.IsErrClientNotFound(err) {
		t.Errorf("Expected Error: 'ErrClientNotFound' but got [%v]", err)
	}
}

func TestChangeClientStatusControllerErrInvalidStatus(t *testing.T) {
	dropCollections()
	/** TO WORK **/
	username := "James Rolfe"
	models.CreateClient(username)
	/**/

	client, err := ChangeClientStatusController(username, "Nerd")

	if client != nil {
		t.Errorf("Expected client to be nil but it was created")
	}
	if !chaterrors.IsErrInvalidStatus(err) {
		t.Errorf("Expected Error: 'ErrInvalidStatus' but got [%v]", err)
	}
}
