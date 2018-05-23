package controllers

import (
	"fmt"
	"log"
	"strings"

	"gitlab.lisbontechhub.com/A-Team/challenger-chat/chaterrors"

	"github.com/googollee/go-socket.io"
)

// SocketResponseFunc that represents the type returned by the socket functions
type SocketResponseFunc func(n map[string]interface{})

// RequestSocketInformation that represents the converted information sent from the client to the socket
type RequestSocketInformation struct {
	Username   string
	FriendName string
	Text       string
	Status     string
}

/**
Public Functions
*/

// SocketInterface that represents the Server side socket.
type SocketInterface interface {
	BroadcastTo(room, message string, args ...interface{})
}

// LoginUserSocketController function that logs the user in the chat application.
func LoginUserSocketController(socket SocketInterface, so socketio.Socket) SocketResponseFunc {
	return func(n map[string]interface{}) {

		convertedData, error := convertMapData(n)
		if error != nil {
			so.Emit("error", error)
			return
		}

		fmt.Printf("\n    --> CONNECTED <--\n")
		fmt.Println("-------------------------")
		fmt.Printf("    username: %v\n", n["username"])
		fmt.Println("-------------------------\n")

		//We start the client
		username := convertedData.Username
		client, err := StartClientController(username)
		if chaterrors.CheckError(err) {
			error := chaterrors.NewError("Start Client Error", err.Error())
			so.Emit("error", error)
			return
		}
		//After we create the session the user will run on
		session, err := StartSessionController(so.Id(), client.Name)
		chaterrors.CheckError(err)

		type UserInfo map[string]interface{}

		userInfo := UserInfo{"Username": username, "SocketId": so.Id(), "SessionId": session.ID}

		so.Emit("user_session", userInfo)
		socket.BroadcastTo("chat", "notify_online", username)
	}
}

// LogoutUserSocketController function that logs out the user from the chat application.
func LogoutUserSocketController(socket SocketInterface, so socketio.Socket) SocketResponseFunc {
	return func(n map[string]interface{}) {
		convertedData, error := convertMapData(n)
		if error != nil {
			so.Emit("error", error)
			return
		}

		username := convertedData.Username

		_, err := LogoutClientController(username)
		if chaterrors.CheckError(err) {
			error := chaterrors.NewError("Logout Client Error", err.Error())
			so.Emit("error", error)
			return
		}
		CloseSessionController(username)
		fmt.Println("*** " + username + " has disconnected ***")

		type UserInfo map[string]interface{}

		userInfo := UserInfo{"User": username, "Status": "Offline"}

		socket.BroadcastTo("chat", "notify_offline", username)
		socket.BroadcastTo("chat", "notify_status", userInfo)
		so.Leave("chat")
		so.Disconnect()
	}
}

// AddFriendSocketController function that adds a friend to the user's list.
func AddFriendSocketController(so socketio.Socket) SocketResponseFunc {
	return func(n map[string]interface{}) {
		convertedData, error := convertMapData(n)
		if error != nil {
			so.Emit("error", error)
			return
		}

		username := convertedData.Username
		friend := convertedData.FriendName

		err := ClientAddFriendController(username, friend)
		if chaterrors.CheckError(err) {
			error := chaterrors.NewError("Add Friend Error", err.Error())
			so.Emit("error", error)
			return
		}
		fmt.Println("*** " + username + " has added " + friend + " as a friend. ***")

		friends, err := GetFriendsController(username)

		clientInfo, err := GetClientController(username)
		chaterrors.CheckError(err)

		so.Emit("get_user_friends", friends)
		so.BroadcastTo("chat", "get_new_friend", clientInfo, friend)
	}
}

// GetChatLogSocketController function that fetches the chat log between two users.
func GetChatLogSocketController(so socketio.Socket) SocketResponseFunc {
	return func(n map[string]interface{}) {
		fmt.Println("--------")
		fmt.Printf("%v+", n)

		convertedData, error := convertMapData(n)
		if error != nil {
			so.Emit("error", error)
			return
		}

		username := convertedData.Username
		friendName := convertedData.FriendName

		messages, err := GetClientChatLogController(username, friendName)
		chaterrors.CheckError(err)

		client, err := GetClientController(username)
		if chaterrors.CheckError(err) {
			error := chaterrors.NewError("Get Client Error", err.Error())
			so.Emit("error", error)
			return
		}

		var roomName string
		for _, v := range client.Friends {
			name := strings.Split(v, "-")
			if name[0] == friendName {
				roomName = name[1]
			}
		}

		so.Join(roomName)
		log.Printf("*** Connected " + username + " and " + friendName + " to room [ " + roomName + " ]")
		so.Emit("get_chat_log", messages)
	}
}

// MessageFriendSocketController function that messages a friend of the user.
func MessageFriendSocketController(so socketio.Socket) SocketResponseFunc {
	return func(n map[string]interface{}) {
		convertedData, error := convertMapData(n)
		if error != nil {
			so.Emit("error", error)
			return
		}

		sender := convertedData.Username
		receiver := convertedData.FriendName
		textMessage := convertedData.Text

		message, err := CreateMessageController(textMessage, sender, receiver)
		if chaterrors.CheckError(err) {
			error := chaterrors.NewError("Send Message Error", err.Error())
			so.Emit("error", error)
			return
		}
		fmt.Printf("%v", message)

		client, err := GetClientController(sender)
		chaterrors.CheckError(err)

		var roomName string
		for _, v := range client.Friends {
			name := strings.Split(v, "-")
			if name[0] == receiver {
				roomName = name[1]
			}
		}
		so.BroadcastTo(roomName, "get_new_message", message)
	}
}

// GetUserFriendsSocketController function that fetches the list of friends of the user.
func GetUserFriendsSocketController(so socketio.Socket) SocketResponseFunc {
	return func(n map[string]interface{}) {
		convertedData, error := convertMapData(n)
		if error != nil {
			so.Emit("error", error)
			return
		}

		username := convertedData.Username
		friends, err := GetFriendsController(username)
		if chaterrors.CheckError(err) {
			error := chaterrors.NewError("Get Friends Error", err.Error())
			so.Emit("error", error)
			return
		}
		fmt.Println("*** Retrieved friends of user [ " + username + " ] ***")
		so.Emit("get_user_friends", friends)
	}
}

// ChangeStatusSocketController function that changes the status of the user.
func ChangeStatusSocketController(so socketio.Socket) SocketResponseFunc {
	return func(n map[string]interface{}) {
		convertedData, error := convertMapData(n)
		if error != nil {
			so.Emit("error", error)
			return
		}

		username := convertedData.Username
		status := convertedData.Status
		_, err := ChangeClientStatusController(username, status)
		chaterrors.CheckError(err)

		client, err := GetClientController(username)
		if chaterrors.CheckError(err) {
			error := chaterrors.NewError("Get Client Error", err.Error())
			so.Emit("error", error)
			return
		}

		//Notify friends that we changed status
		for _, v := range client.Friends {

			friendName := strings.Split(v, "-")

			session, err := GetUserSessionController(friendName[0])
			if chaterrors.CheckError(err) {
				error := chaterrors.NewError("Get User Session", err.Error())
				so.Emit("error", error)
				return
			}

			if session.SocketID != "" {
				type UserInfo map[string]interface{}

				userInfo := UserInfo{"User": username, "Status": status}

				so.BroadcastTo("chat", "notify_status", userInfo)
			}
		}
	}
}

// CloseUserSessionSocketController function that disconnects the user from the chat application.
func CloseUserSessionSocketController(so socketio.Socket) SocketResponseFunc {
	return func(n map[string]interface{}) {
		fmt.Printf("\n   --> DISCONNECTED %v <--\n ", so.Id())
		DisconnectUserSessionController(so.Id())
	}
}

/**
Private Functions
*/

// convertMapData function that receives a map of strings and convert it
// into a RequestSocketInformation object to be read by the socket.
func convertMapData(n map[string]interface{}) (*RequestSocketInformation, *chaterrors.SocketError) {
	data := new(RequestSocketInformation)

	for k, v := range n {
		switch k {
		case "username":
			if chaterrors.CheckIncorrectString(v) {
				error := chaterrors.NewError("Incorrect Name Format", "The username has an incorrect format.")
				return nil, &error
			}
			data.Username = v.(string)
		case "friendname":
			if chaterrors.CheckIncorrectString(v) {
				error := chaterrors.NewError("Incorrect Friendname Format", "The friendname has an incorrect format.")
				return nil, &error
			}
			data.FriendName = v.(string)
		case "text":
			if chaterrors.CheckIncorrectString(v) {
				error := chaterrors.NewError("Incorrect Text Format", "The text has an incorrect format.")
				return nil, &error
			}
			data.Text = v.(string)
		case "status":
			if chaterrors.CheckIncorrectString(v) {
				error := chaterrors.NewError("Incorrect Status Format", "The status has an incorrect format.")
				return nil, &error
			}
			data.Status = v.(string)
		default:
			if chaterrors.CheckIncorrectString(v) {
				error := chaterrors.NewError("Wrong Field Error", "This field is not valid.")
				return nil, &error
			}
		}
	}
	return data, nil
}
