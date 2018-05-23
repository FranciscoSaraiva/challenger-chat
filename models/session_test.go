package models

import (
	"testing"
)

func TestCreateSessionCorrect(t *testing.T) {
	socketID := "socketIDtest123"
	testUser := "TestUser"
	session, err := CreateSession(socketID, testUser)
	if checkError(err) {
		t.Errorf("Expected no error and got [%v]", err)
	}
	if session.SocketID != socketID {
		t.Errorf("Expected SocketID: [%s] got [%s]", socketID, session.SocketID)
	}
	if session.User != testUser {
		t.Errorf("Expected User: [%s] got [%s]", testUser, session.User)
	}
}

func TestUpdateSessionCorrect(t *testing.T) {
	testUser := "TestUser"
	oldSocketID := "socketIDtest123"
	CreateSession(oldSocketID, testUser)

	newSocketID := "NewSocketIDtest123"

	session, err := UpdateSession(newSocketID, testUser)
	if checkError(err) {
		t.Errorf("Expected no error and got [%v]", err)
	}
	if session.SocketID == oldSocketID {
		t.Errorf("Expected SocketID to update to [%s] but is still [%s]", newSocketID, oldSocketID)
	}
}

func TestCloseSessionCorrect(t *testing.T) {
	user := "TestUser"
	session := CloseSession(user)

	if session.Active == true {
		t.Errorf("Expected Session to close but value is [%v]", session.Active)
	}
}

func TestGetUserSessionCorrect(t *testing.T) {
	user := "TestUser"
	session, err := GetUserSession(user)
	if checkError(err) {
		t.Errorf("Expected no error and got [%v]", err)
	}
	if session.User != user {
		t.Errorf("Expected User: [%s] but got [%s]", user, session.User)
	}
}

func TestDisconnectUserSessionCorrect(t *testing.T) {
	socketID := "NewSocketIDtest123"
	session := DisconnectUserSession(socketID)
	if session.Active == true {
		t.Errorf("Expected Session to close but value is [%v]", session.Active)
	}
}
