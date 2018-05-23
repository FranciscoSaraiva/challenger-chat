package controllers

import (
	"reflect"
	"testing"

	"gitlab.lisbontechhub.com/A-Team/challenger-chat/chaterrors"
	"gitlab.lisbontechhub.com/A-Team/challenger-chat/models"
)

func TestStartSessionControllerCreate(t *testing.T) {
	dropCollections()
	/** FOR IT TO WORK **/
	username := "James Rolfe"
	client, _ := models.CreateClient(username)
	socketID := "SockedIDTest123"
	/**/

	session, err := StartSessionController(socketID, client.Name)

	if checkError(err) {
		t.Errorf("Expected no error but got Error: [%v]", err)
	}
	if session.Active == false {
		t.Errorf("Expected Session to be true but it was [%v]", session.Active)
	}
	if session.User != client.Name {
		t.Errorf("Expected User [%s] but got [%s]", client.Name, session.User)
	}
	if session.SocketID != socketID {
		t.Errorf("Expected SocketID [%s] but got the id [%s]", socketID, session.SocketID)
	}
	if session.Time == 0 {
		t.Errorf("Expected a Timestamp to exist but it is 0")
	}
}

func TestStartSessionControllerUpdate(t *testing.T) {
	dropCollections()
	/** FOR IT TO WORK **/
	username := "James Rolfe"
	client, _ := models.CreateClient(username)
	socketID := "SockedIDTest123"
	models.CreateSession(socketID, client.Name)
	/**/

	session, err := StartSessionController(socketID, client.Name)

	if checkError(err) {
		t.Errorf("Expected no error but got Error: [%v]", err)
	}
	if session.Active == false {
		t.Errorf("Expected Session to be true but it was [%v]", session.Active)
	}
	if session.User != client.Name {
		t.Errorf("Expected User [%s] but got [%s]", client.Name, session.User)
	}
	if session.SocketID != socketID {
		t.Errorf("Expected SocketID [%s] but got the id [%s]", socketID, session.SocketID)
	}
	if session.Time == 0 {
		t.Errorf("Expected a Timestamp to exist but it is 0")
	}
}

func TestStartSessionControllerErrClientNotFound(t *testing.T) {
	dropCollections()
	/** FOR IT TO WORK **/
	username := "James Rolfe"
	socketID := "SockedIDTest123"
	/**/

	session, err := StartSessionController(socketID, username)

	if session != nil {
		t.Errorf("Expected session to not have been created but it's not nil")
	}
	if !chaterrors.IsErrClientNotFound(err) {
		t.Errorf("Expected error to be of type 'ErrClientNotFound' but got Error: [%v]", reflect.TypeOf(err))
	}
}

func TestCloseSessionControllerCorrect(t *testing.T) {
	dropCollections()
	/** FOR IT TO WORK **/
	username := "James Rolfe"
	client, _ := models.CreateClient(username)
	socketID := "SockedIDTest123"
	/**/

	StartSessionController(socketID, client.Name)
	session, err := CloseSessionController(client.Name)

	if checkError(err) {
		t.Errorf("Expected no error but got [%v]", err)
	}
	if session.Active == true {
		t.Errorf("Expected Session to be closed but is active")
	}
}

func TestCloseSessionControllerErrSessionNotFound(t *testing.T) {
	dropCollections()
	/** FOR IT TO WORK **/
	username := "James Rolfe"
	client, _ := models.CreateClient(username)
	/**/

	session, err := CloseSessionController(client.Name)

	if session != nil {
		t.Errorf("Expected Session to not be created but it was")
	}
	if !chaterrors.IsErrSessionNotFound(err) {
		t.Errorf("Expected error to be of type 'ErrSessionNotFound' but got Error: [%v]", reflect.TypeOf(err))
	}
}

func TestGetUserSessionControllerCorrect(t *testing.T) {
	dropCollections()
	/** FOR IT TO WORK **/
	username := "James Rolfe"
	client, _ := models.CreateClient(username)
	socketID := "SockedIDTest123"
	/**/

	session, _ := StartSessionController(socketID, client.Name)
	sessionGotten, err := GetUserSessionController(client.Name)

	if checkError(err) {
		t.Errorf("Expected no error but got [%v]", err)
	}
	if sessionGotten.SocketID != session.SocketID {
		t.Errorf("Expected Session to have SocketID [%s] but was found [%s]", session.SocketID, sessionGotten.SocketID)
	}
	if session.User != sessionGotten.User {
		t.Errorf("Expected Session to belong to [%s] but was found [%s]", session.User, sessionGotten.User)
	}
}

func TestGetUserSessionControllerErrSessionNotFound(t *testing.T) {
	dropCollections()
	/** FOR IT TO WORK **/
	username := "James Rolfe"
	client, _ := models.CreateClient(username)
	//socketID := "SockedIDTest123"
	/**/

	session, err := GetUserSessionController(client.Name)

	if session != nil {
		t.Errorf("Expected Session to be nil but it was  created")
	}
	if !chaterrors.IsErrSessionNotFound(err) {
		t.Errorf("Expected error to be of type 'ErrSessionNotFound' but got Error: [%v]", reflect.TypeOf(err))
	}
}

func TestGetUserSessionControllerErrClientNotFound(t *testing.T) {
	dropCollections()
	/** FOR IT TO WORK **/
	username := "James Rolfe"
	socketID := "SockedIDTest123"
	/**/

	models.CreateSession(socketID, username)
	sessionGotten, err := GetUserSessionController(username)

	if sessionGotten != nil {
		t.Errorf("Expected Session to be nil but it was  created")
	}
	if !chaterrors.IsErrClientNotFound(err) {
		t.Errorf("Expected error to be of type 'ErrClientNotFound' but got Error: [%v]", reflect.TypeOf(err))
	}
}

func TestDisconnectUserSessionControllerCorrect(t *testing.T) {
	dropCollections()
	/** FOR IT TO WORK **/
	username := "James Rolfe"
	client, _ := models.CreateClient(username)
	socketID := "SockedIDTest123"
	/**/

	session, _ := StartSessionController(socketID, client.Name)
	sessionD := DisconnectUserSessionController(session.SocketID)

	if sessionD.Active == true {
		t.Errorf("Expected Session to be inactive but it is still active")
	}
}
