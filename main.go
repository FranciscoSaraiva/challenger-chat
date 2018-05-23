package main

import (
	"fmt"
	"log"
	"net/http"
	"challenger-chat/chaterrors"
	"challenger-chat/controllers"

	"github.com/googollee/go-socket.io"
)

//The main socket that will run on the server
var server *socketio.Server

func main() {
	// Initialization
	var err error

	//Connections
	server = getSocket()
	chaterrors.CheckError(err)

	//Endpoints
	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./www")))
	//Console info
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

// getSocket function that returns a new socket.
// It is used to attach it to the main global Socket and keep the work
// environment clean.
func getSocket() *socketio.Server {
	//Socket initialization
	socket, err := socketio.NewServer(nil)
	chaterrors.CheckError(err)

	//Socket configuration
	socket.On("connection", func(so socketio.Socket) {
		fmt.Printf("\n   --> CONNECTED <--\n")

		so.Join("chat") //join the room
		fmt.Println("Joined room [chat]")

		so.On("login_user", controllers.LoginUserSocketController(socket, so))

		so.On("logout_user", controllers.LogoutUserSocketController(socket, so))

		so.On("add_friend", controllers.AddFriendSocketController(so))

		so.On("get_chat_log", controllers.GetChatLogSocketController(so))

		so.On("message_friend", controllers.MessageFriendSocketController(so))

		so.On("get_user_friends", controllers.GetUserFriendsSocketController(so))

		so.On("change_status", controllers.ChangeStatusSocketController(so))

		so.On("disconnection", controllers.CloseUserSessionSocketController(so))
	})

	socket.On("error", func(so socketio.Socket, err error) {
		error := chaterrors.NewError("Socket Error", err.Error())
		so.Emit("error", error)
		return
	})
	return socket
}
