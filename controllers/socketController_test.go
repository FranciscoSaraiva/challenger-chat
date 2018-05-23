package controllers

import (
	"errors"
	"net/http"
	"testing"
)

type mockSocketServer struct {
	IsBroadcastToCalled bool
}

func (s *mockSocketServer) BroadcastTo(room, message string, args ...interface{}) {
	s.IsBroadcastToCalled = true
}

type mockSocketClient struct {
	IsEmitCalled        bool
	IsBroadcastToCalled bool
}

func (s *mockSocketClient) Emit(event string, args ...interface{}) error {
	s.IsEmitCalled = true
	return nil
}

type UserInfo map[string]interface{}

func (s mockSocketClient) Id() string                           { return "" }
func (s mockSocketClient) Rooms() []string                      { return []string{""} }
func (s mockSocketClient) Request() *http.Request               { return nil }
func (s mockSocketClient) On(event string, f interface{}) error { return errors.New("") }
func (s mockSocketClient) Join(room string) error               { return errors.New("") }
func (s mockSocketClient) Leave(room string) error              { return errors.New("") }
func (s mockSocketClient) Disconnect()                          {}
func (s *mockSocketClient) BroadcastTo(room, event string, args ...interface{}) error {
	s.IsBroadcastToCalled = true
	return nil
}

/*****
** LoginUserSocketController
******/

func TestLoginUserSocketControllerCorrect(t *testing.T) {
	dropCollections()
	server := new(mockSocketServer)
	client := new(mockSocketClient)

	callback := LoginUserSocketController(server, client)
	username := "James Rolfe"
	params := UserInfo{"username": username}
	callback(params)

	if server.IsBroadcastToCalled == false {
		t.Errorf("Expected server to broadcast but it didn't")
	}
	if client.IsEmitCalled == false {
		t.Errorf("Expected client to emit but it didn't")
	}
}

func TestLoginUserSocketControllerErrConvert(t *testing.T) {
	dropCollections()
	server := new(mockSocketServer)
	client := new(mockSocketClient)

	callback := LoginUserSocketController(server, client)

	username := 12
	params := UserInfo{"username": username}
	callback(params)

	if client.IsEmitCalled == false {
		t.Errorf("Expected client to emit but it didn't")
	}
}

func TestLoginUserSocketControllerErrClient(t *testing.T) {
	dropCollections()
	server := new(mockSocketServer)
	client := new(mockSocketClient)

	callback := LoginUserSocketController(server, client)

	username := ""
	params := UserInfo{"username": username}
	callback(params)

	callbackAgain := LoginUserSocketController(server, client)

	callbackAgain(params)

	if client.IsEmitCalled == false {
		t.Errorf("Expected client to emit but it didn't")
	}
}

/*****
** LogoutUserSocketController
******/

func TestLogoutUserSocketControllerCorrect(t *testing.T) {
	dropCollections()
	server := new(mockSocketServer)
	client := new(mockSocketClient)

	callback := LogoutUserSocketController(server, client)

	username := "James Rolfe"
	StartClientController(username)
	params := UserInfo{"username": username}
	callback(params)

	if server.IsBroadcastToCalled == false {
		t.Errorf("Expected server to broadcast but it didn't")
	}
}

func TestLogoutUserSocketControllerErrConvert(t *testing.T) {
	dropCollections()
	server := new(mockSocketServer)
	client := new(mockSocketClient)

	callback := LogoutUserSocketController(server, client)

	username := 12
	params := UserInfo{"username": username}
	callback(params)

	if client.IsEmitCalled == false {
		t.Errorf("Expected the emit to be called but it was not.")
	}
}

func TestLogoutUserSocketControllerErrLogout(t *testing.T) {
	dropCollections()
	server := new(mockSocketServer)
	client := new(mockSocketClient)

	callback := LogoutUserSocketController(server, client)

	username := "James Rolfe"
	params := UserInfo{"username": username}
	callback(params)

	if client.IsEmitCalled == false {
		t.Errorf("Expected the emit to be called but it was not.")
	}
}

/*****
** AddFriendSocketController
******/

func TestAddFriendSocketControllerCorrect(t *testing.T) {
	dropCollections()
	client := new(mockSocketClient)

	username := "James Rolfe"
	StartClientController(username)
	friendname := "John Doe"
	StartClientController(friendname)
	params := UserInfo{"username": username, "friendname": friendname}

	callback := AddFriendSocketController(client)
	callback(params)

	if client.IsEmitCalled == false {
		t.Errorf("Expected client to emit but it didn't")
	}
	if client.IsBroadcastToCalled == false {
		t.Errorf("Expected server to broadcast but it didn't")
	}
}

func TestAddFriendSocketControllerErrConvert(t *testing.T) {
	client := new(mockSocketClient)

	callback := AddFriendSocketController(client)

	username := "James Rolfe"
	friendname := 12
	params := UserInfo{"username": username, "friendname": friendname}
	callback(params)

	if client.IsEmitCalled == false {
		t.Errorf("Expected the emit to be called but it was not.")
	}
}

func TestAddFriendSocketControllerErrAddFriend(t *testing.T) {
	client := new(mockSocketClient)

	callback := AddFriendSocketController(client)

	username := "James Rolfe"
	friendname := "John Doe"
	params := UserInfo{"username": username, "friendname": friendname}
	callback(params)

	if client.IsEmitCalled == false {
		t.Errorf("Expected client to emit but it didn't")
	}
}

func TestAddFriendSocketControllerErrConvertInvalid(t *testing.T) {
	dropCollections()
	client := new(mockSocketClient)

	username := "James Rolfe"
	StartClientController(username)
	friendname := "John Doe"
	StartClientController(friendname)
	params := UserInfo{"incorrect": 12, "invalid": 12}

	callback := AddFriendSocketController(client)
	callback(params)
	if client.IsEmitCalled == false {
		t.Errorf("Expected client to emit but it didn't")
	}
}

/*****
** GetChatLogSocketController
******/

func TestGetChatLogSocketControllerCorrect(t *testing.T) {
	dropCollections()
	client := new(mockSocketClient)

	callback := GetChatLogSocketController(client)

	username := "James Rolfe"
	StartClientController(username)
	friendname := "John Doe"
	StartClientController(friendname)
	ClientAddFriendController(username, friendname)
	text := "Oi"
	CreateMessageController(text, username, friendname)
	CreateMessageController(text, username, friendname)
	CreateMessageController(text, username, friendname)

	params := UserInfo{"username": username, "friendname": friendname, "text": text}
	callback(params)

	if client.IsEmitCalled == false {
		t.Errorf("Expected client to emit but it didn't")
	}
}

func TestGetChatLogSocketControllerErrConvert(t *testing.T) {
	dropCollections()
	client := new(mockSocketClient)

	callback := GetChatLogSocketController(client)

	username := "James Rolfe"
	StartClientController(username)
	friendname := "John Doe"
	StartClientController(friendname)
	text := "Oi"
	CreateMessageController(text, username, friendname)
	CreateMessageController(text, username, friendname)
	CreateMessageController(text, username, friendname)

	params := UserInfo{"username": username, "friendname": friendname, "text": 12}
	callback(params)
	if client.IsEmitCalled == false {
		t.Errorf("Expected client to emit but it didn't")
	}
}

func TestGetChatLogSocketControllerErrChatLog(t *testing.T) {
	dropCollections()
	client := new(mockSocketClient)

	callback := GetChatLogSocketController(client)

	username := "James Rolfe"
	friendname := "John Doe"
	text := "Oi"
	CreateMessageController(text, username, friendname)
	CreateMessageController(text, username, friendname)
	CreateMessageController(text, username, friendname)

	params := UserInfo{"username": username, "friendname": friendname, "text": text}
	callback(params)

	if client.IsEmitCalled == false {
		t.Errorf("Expected client to emit but it didn't")
	}
}

/*****
** MessageFriendSocketController
******/

func TestMessageFriendSocketControllerCorrect(t *testing.T) {
	dropCollections()
	client := new(mockSocketClient)

	callback := MessageFriendSocketController(client)

	username := "James Rolfe"
	StartClientController(username)
	friendname := "John Doe"
	StartClientController(friendname)
	ClientAddFriendController(username, friendname)
	params := UserInfo{"username": username, "friendname": friendname, "text": "Oii"}
	callback(params)

	if client.IsBroadcastToCalled == false {
		t.Errorf("Expected server to broadcast but it didn't")
	}
}

func TestMessageFriendSocketControllerErrConv(t *testing.T) {
	dropCollections()
	client := new(mockSocketClient)

	callback := MessageFriendSocketController(client)

	username := "James Rolfe"
	StartClientController(username)
	friendname := "John Doe"
	StartClientController(friendname)
	params := UserInfo{"username": 12, "friendname": friendname, "text": "Oii"}
	callback(params)

	if client.IsEmitCalled == false {
		t.Errorf("Expected client to emit but it didn't")
	}
}

func TestMessageFriendSocketControllerErrMessage(t *testing.T) {
	dropCollections()
	client := new(mockSocketClient)

	callback := MessageFriendSocketController(client)

	username := "James Rolfe"
	friendname := "John Doe"
	params := UserInfo{"username": username, "friendname": friendname, "text": "Oii"}
	callback(params)

	if client.IsEmitCalled == false {
		t.Errorf("Expected client to emit but it didn't")
	}
}

/*****
** GetUserFriendsSocketController
******/

func TestGetUserFriendsSocketControllerCorrect(t *testing.T) {
	dropCollections()
	client := new(mockSocketClient)

	callback := GetUserFriendsSocketController(client)

	username := "James Rolfe"
	StartClientController(username)
	friendname := "John Doe"
	StartClientController(friendname)
	ClientAddFriendController(username, friendname)
	params := UserInfo{"username": username}
	callback(params)

	if client.IsEmitCalled == false {
		t.Errorf("Expected client to emit but it didn't")
	}
}

func TestGetUserFriendsSocketControllerErrConv(t *testing.T) {
	dropCollections()
	client := new(mockSocketClient)

	callback := GetUserFriendsSocketController(client)

	username := "James Rolfe"
	StartClientController(username)
	friendname := "John Doe"
	StartClientController(friendname)
	ClientAddFriendController(username, friendname)
	params := UserInfo{"username": 12}
	callback(params)

	if client.IsEmitCalled == false {
		t.Errorf("Expected client to emit but it didn't")
	}
}

func TestGetUserFriendsSocketControllerErrGetFriends(t *testing.T) {
	dropCollections()
	client := new(mockSocketClient)

	callback := GetUserFriendsSocketController(client)

	username := "James Rolfe"
	params := UserInfo{"username": username}
	callback(params)

	if client.IsEmitCalled == false {
		t.Errorf("Expected client to emit but it didn't")
	}
}

/*****
** ChangeStatusSocketController
******/

func TestChangeStatusSocketControllerCorrect(t *testing.T) {
	dropCollections()
	client := new(mockSocketClient)

	callback := ChangeStatusSocketController(client)

	username := "James Rolfe"
	StartClientController(username)
	StartSessionController("JamesSocketID", username)
	friendname := "John Doe"
	StartClientController(friendname)
	StartSessionController("JohnSocketID", friendname)
	ClientAddFriendController(username, friendname)
	status := "Busy"
	params := UserInfo{"username": username, "status": status}
	callback(params)

	if client.IsBroadcastToCalled == false {
		t.Errorf("Expected server to broadcast but it didn't")
	}
}

func TestChangeStatusSocketControllerErrConv(t *testing.T) {
	dropCollections()
	client := new(mockSocketClient)

	callback := ChangeStatusSocketController(client)

	username := "James Rolfe"
	StartClientController(username)
	friendname := "John Doe"
	StartClientController(friendname)
	ClientAddFriendController(username, friendname)
	status := 12
	params := UserInfo{"username": username, "status": status}
	callback(params)

	if client.IsEmitCalled == false {
		t.Errorf("Expected client to emit but it didn't")
	}
}

func TestChangeStatusSocketControllerErrChangeStatus(t *testing.T) {
	dropCollections()
	client := new(mockSocketClient)

	callback := ChangeStatusSocketController(client)

	username := "James Rolfe"
	status := "Busy"
	params := UserInfo{"username": username, "status": status}
	callback(params)

	if client.IsEmitCalled == false {
		t.Errorf("Expected client to emit but it didn't")
	}
}

func TestChangeStatusSocketControllerErrSocketNotEmpty(t *testing.T) {
	dropCollections()
	client := new(mockSocketClient)

	callback := ChangeStatusSocketController(client)

	username := "James Rolfe"
	StartClientController(username)
	StartSessionController("JamesSocketID", username)
	friendname := "John Doe"
	StartClientController(friendname)
	StartSessionController("JohnSocketID", friendname)

	ClientAddFriendController(username, friendname)

	status := "Busy"
	params := UserInfo{"username": username, "status": status}
	callback(params)

	if client.IsBroadcastToCalled == false {
		t.Errorf("Expected client to emit but it didn't")
	}
}

func TestChangeStatusSocketControllerErrGetUserSession(t *testing.T) {
	dropCollections()
	client := new(mockSocketClient)

	callback := ChangeStatusSocketController(client)

	username := "James Rolfe"
	StartClientController(username)
	friendname := "John Doe"
	StartClientController(friendname)
	ClientAddFriendController(username, friendname)
	status := "Busy"
	params := UserInfo{"username": username, "status": status}
	callback(params)

	if client.IsEmitCalled == false {
		t.Errorf("Expected client to emit but it didn't")
	}
}

/*****
** CloseUserSessionSocketController
******/

func TestCloseUserSessionSocketController(t *testing.T) {
	client := new(mockSocketClient)

	callback := CloseUserSessionSocketController(client)

	username := "James Rolfe"
	params := UserInfo{"username": username}
	callback(params)
}
