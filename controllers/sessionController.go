package controllers

import (
	"gitlab.lisbontechhub.com/A-Team/challenger-chat/chaterrors"
	"gitlab.lisbontechhub.com/A-Team/challenger-chat/config"
	"gitlab.lisbontechhub.com/A-Team/challenger-chat/models"

	"gopkg.in/mgo.v2/bson"
)

/*
 Public Functions
*/

// StartSessionController function that redirects what action to do based on the socketID and the username.
func StartSessionController(socketID, user string) (*models.Session, error) {
	if !CheckIfUserExists(user) {
		return nil, chaterrors.ErrClientNotFound{Username: user}
	}
	if checkIfSessionExists(user) {
		session, err := models.UpdateSession(socketID, user)
		return session, err
	}
	session, err := models.CreateSession(socketID, user)
	return session, err
}

// CloseSessionController function that controls the flow to close a session of a user.
func CloseSessionController(user string) (*models.Session, error) {
	if !checkIfSessionExists(user) {
		return nil, chaterrors.ErrSessionNotFound{Username: user}
	}
	session := models.CloseSession(user)
	return session, nil
}

// GetUserSessionController function that controls the flow to getting a session of a user.
func GetUserSessionController(user string) (*models.Session, error) {
	if !checkIfSessionExists(user) {
		return nil, chaterrors.ErrSessionNotFound{Username: user}
	}
	if !CheckIfUserExists(user) {
		return nil, chaterrors.ErrClientNotFound{Username: user}
	}
	session, err := models.GetUserSession(user)
	return session, err
}

// DisconnectUserSessionController function that controls the flow to disconnecting the user's session.
func DisconnectUserSessionController(socketID string) *models.Session {
	session := models.DisconnectUserSession(socketID)
	return session
}

/*
 Private Functions
*/

// checkIfSessionExists function that checks if a user is currently connected in an existing
// session in the Socket.
func checkIfSessionExists(user string) bool {
	sessions := config.GetCollection("Session")
	count, _ := sessions.Find(bson.M{"user": user}).Count()
	if count == 1 { //meaning one was found
		return true
	}
	return false
}
