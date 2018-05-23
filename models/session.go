package models

import (
	"fmt"

	"gitlab.lisbontechhub.com/A-Team/challenger-chat/chaterrors"
	"gitlab.lisbontechhub.com/A-Team/challenger-chat/config"

	"gopkg.in/mgo.v2/bson"
)

// Session structure that represents the socket session a user is connected to.
// Saves the SocketID the user is connected to, Active represents if he's currently
// connected, Time shows the time of connection, User is the name of the User connected.
type Session struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	SocketID string
	Active   bool
	Time     int64
	User     string
}

/*
 Public Functions
*/

// CreateSession function that creates a new Session with information of the user's session in the
// the Socket he's connected in.
func CreateSession(socketID, user string) (*Session, error) {
	session := new(Session)
	session.ID = bson.NewObjectId()
	session.SocketID = socketID
	session.Active = true
	session.Time = config.GetTimeNowTimeStamp()
	session.User = user

	collection := config.GetCollection("Session")
	err := collection.Insert(&session)

	fmt.Printf(" %+v", session)
	return session, err
}

// UpdateSession function that updates the information of the Session the user is currently connected
// to in the Socket.
func UpdateSession(socketID, user string) (*Session, error) {
	collection := config.GetCollection("Session")
	query := bson.M{"user": user}

	session := new(Session)
	err := collection.Find(query).One(session)
	chaterrors.CheckError(err)

	session.SocketID = socketID
	session.Active = true
	session.Time = config.GetTimeNowTimeStamp()
	collection.Update(query, session)

	return session, err
}

// CloseSession function that closes the session a user was connected to, via Socket.
func CloseSession(user string) *Session {
	collection := config.GetCollection("Session")
	query := bson.M{"user": user}

	session := new(Session)
	err := collection.Find(query).One(session)
	chaterrors.CheckError(err)

	session.SocketID = ""
	session.Active = false
	session.Time = config.GetTimeNowTimeStamp()
	collection.Update(query, session)

	return session
}

// GetUserSession function that returned an object of the Session with the information about
// a user's session in the Socket.
func GetUserSession(user string) (*Session, error) {
	collection := config.GetCollection("Session")
	query := bson.M{"user": user}

	session := new(Session)
	err := collection.Find(query).One(session)

	return session, err
}

// DisconnectUserSession function that disconnects the user from the session is he on
// in the case he disconnects without choosing the logout option.
func DisconnectUserSession(socketID string) *Session {
	collection := config.GetCollection("Session")
	query := bson.M{"socketid": socketID}

	session := new(Session)
	err := collection.Find(query).One(session)
	chaterrors.CheckError(err)
	session.SocketID = ""
	session.Active = false
	session.Time = config.GetTimeNowTimeStamp()
	collection.Update(query, session)

	collection = config.GetCollection("Client")
	query = bson.M{"name": session.User}

	client := new(Client)
	err = collection.Find(query).One(client)
	chaterrors.CheckError(err)
	client.Status = "Offline"
	collection.Update(query, client)

	return session
}
