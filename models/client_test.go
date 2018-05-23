package models

import (
	"log"
	"strings"
	"testing"
)

func TestCreateClientCorrect(t *testing.T) {
	user := "TestUser"
	client, err := CreateClient(user)
	if checkError(err) {
		t.Errorf("Expected no error and got [%v]", err)
	}
	if client.Name != user {
		t.Errorf("Expected Name: [%s] but got [%s]", user, client.Name)
	}
	if client.Status != "Online" {
		t.Errorf("Expected Status to be 'Online but got [%s]", client.Status)
	}
	if len(client.Friends) != 0 {
		t.Errorf("Expected Friends count to be 0 but got [%d]", len(client.Friends))
	}
}

func TestLoginClientCorrect(t *testing.T) {
	user := "TestUser"
	ClientAddFriend(user, "TestUserTwo")
	client, err := LoginClient(user)
	if checkError(err) {
		t.Errorf("Expected no error and got [%v]", err)
	}
	if client.Status != "Online" {
		t.Errorf("Expected the Status to be 'Online' and got [%s]", client.Status)
	}
}

func TestLogoutClientCorrect(t *testing.T) {
	user := "TestUser"
	client, err := LogoutClient(user)
	if checkError(err) {
		t.Errorf("Expected no error and got [%v]", err)
	}
	if client.Status != "Offline" {
		t.Errorf("Expected the Status to be 'Offline' and got [%s]", client.Status)
	}
}

func TestGetClientCorrect(t *testing.T) {
	user := "TestUser"
	client, err := GetClient(user)
	if checkError(err) {
		t.Errorf("Expected no error and got [%v]", err)
	}
	if client.Name != user {
		t.Errorf("Expected Name [%s] and got [%s]", user, client.Name)
	}
}

func TestChangeClientStatusCorrect(t *testing.T) {
	user := "TestUser"
	status := "Away"
	client, err := ChangeClientStatus(user, status)
	if checkError(err) {
		t.Errorf("Expected no error and got [%v]", err)
	}
	if client.Name != user {
		t.Errorf("Expected Name [%s] and got [%s]", user, client.Name)
	}
	if client.Status != "Away" {
		t.Errorf("Expected Status [%s] and got [%s]", status, client.Status)
	}
}

func TestClientAddFriendCorrect(t *testing.T) {
	user := "TestUser"
	friend := "TestUserTwo"
	clientF, _ := CreateClient("TestUserTwo")
	client, err := ClientAddFriend(user, clientF.Name)
	if checkError(err) {
		t.Errorf("Expected no error and got [%v]", err)
	}
	if client.Name != clientF.Name {
		t.Errorf("Expected Friend Name [%s] and got [%s]", friend, client.Name)
	}
	if len(client.Friends) == 0 {
		t.Errorf("Expected friend list to have friends but it's empty")
	}

	for _, v := range client.Friends {
		friendname := strings.Split(v, "-")
		if friendname[0] == user {
			return
		}
	}
	t.Errorf("Expected Friend [%s] to have added Client [%s] but wasn't found", friend, client.Name)
}

func TestGetFriendsCorrect(t *testing.T) {
	user := "TestUser"
	friend := "TestUserTwo"
	friends, err := GetFriends(user)
	if checkError(err) {
		t.Errorf("Expected no error and got [%v]", err)
	}
	if len(friends) == 0 {
		t.Errorf("Expected some friends but got 0 friends")
	}
	if friends[1].Name != friend {
		t.Errorf("Expected Client [%s] to have Friend [%s] as friend but wasn't found \n [%+v]", user, friend, friends[0].Name)
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
