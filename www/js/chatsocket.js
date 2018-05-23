/***************
 *** -> VUE
 ****************/

(function () {
    var username;
    window.app_chat = new Vue({
        el: '#app-chat',
        data: {
            username: "",
            status: "Online",
            friends: [],
            //---
            isUserLogged: false,
            isChatSelected: false,
            statusOptions: [
                "Online",
                "Away",
                "Busy"
            ],
            currentChat: {
                friend_name: null,
                friend_status: null,
                messages: []
            }
        },
        methods: {
            //Login
            loginClient: function(){
                this.status = "Online";
                this.username = prompt('Please enter your username: ');
                username = this.username;
                socket.emit('login_user', {
                    username: this.username
                });
                this.friends = [];
                socket.emit('get_user_friends', {
                    username: this.username
                });
                this.isUserLogged = true;
                this.isChatSelected=false;
                this.changeStatus();
            },

            //Set the list of friends
            setUserFriends: function(friends){
                this.friends = friends;
            },

            //Logout
            logoutClient: function (){
                socket.emit('logout_user', {
                    username: this.username
                });
                this.status="Offline";
                this.changeStatus();
                this.isUserLogged = false;
                this.currentChat.friend_name = "";
                this.currentChat.friend_status = "";
                this.currentChat.messages = [];
            },

            //Add Friend
            addFriend: function (){
                var friend = prompt('Enter the name of the person you want to add')
                socket.emit('add_friend', {
                    username: this.username,
                    friendname: friend
                })
            },

            //Message friend
            sendMessage: function(){
                socket.emit('message_friend', {
                    username: this.username,
                    friendname: this.currentChat.friend_name,
                    text: document.getElementsByClassName('message-text')[0].value
                });

                //pushing the message
                this.currentChat.messages.push({
                    sender: this.username,
                    receiver: this.currentChat.friend_name,
                    text: document.getElementsByClassName('message-text')[0].value,
                    timestamp: + new Date()
                });

                document.getElementsByClassName('message-text')[0].value = "";

                //auto scroll
                //autoScrollChatLog()
            },

            //Change Status
            changeStatus: function(){
                socket.emit('change_status', {
                    username: this.username,
                    status: this.status
                })
            },

            //Get the list of friends
            getUserFriends: function(){
                socket.emit('get_user_friends', {
                    username: this.username
                })
            },

            //Open chat window of friend
            openChat: function(friend_name, friend_status){
                this.isChatSelected=true;
                this.currentChat.friend_name=friend_name;
                this.currentChat.status=friend_status;
                this.currentChat.messages = [];

                //change the current chat
                var td = document.getElementById('friend-bar')
                    .getElementsByTagName('td');
                var p = td[0].getElementsByTagName('p');

                p[0].innerText=friend_name;
                p[1].innerText=friend_status;

                this.changeStatus();

                socket.emit('get_chat_log', {
                    username: this.username,
                    friendname: friend_name
                });
            }
        }
    });

    /***************
     *** -> SOCKET FUNCTIONS
     ****************/
    var socket = io();
    socket.on('connect', function(){
        //Get friend information of the logged in user
        socket.on('get_user_friends', function(data){
            var friends = data.map(mapUsers);
            app_chat.setUserFriends(friends);
        });

        // 0-username, 1-status, 2-socket id of friend
        // Notify of status changes
        socket.on('notify_status', function(data){
            console.group(data)

            //change the friend list entry
            var friend = document.getElementById(data.User);
            var td = friend.getElementsByTagName("td");
            td[1].innerText = data.User+" - "+data.Status;

            //change the current chat
            var td = document.getElementById('friend-bar')
                .getElementsByTagName('td');
            var p = td[0].getElementsByTagName('p');

            if(p[0].innerText === data[0]){ //So that we don't change the current chat window info
                p[1].innerText=data[1];
            }
        });

        //Get the chat log between a friend
        socket.on('get_chat_log', function(data){
            var messages = data.map(mapMessages);
            app_chat.currentChat.messages=messages;
            console.log(messages);
        });

        //On a new message we might get on the chat window
        socket.on('get_new_message', function(data){
            console.log(data, username);
            if (app_chat.currentChat.friend_name === data.Sender){
                app_chat.currentChat.messages.push({
                    sender: data.Sender,
                    receiver: data.Receiver,
                    text: data.Text,
                    timestamp: data.Timestamp
                });

                // auto scroll
                autoScrollChatLog()
            }
        });

        //Get a friend added to our list who adds us
        socket.on('get_new_friend', function(data, clientName){
            if(app_chat.username === clientName){
                app_chat.friends.push({username: data.Name, status: data.Status})
            }
        });

        //Get login information
        socket.on('user_session', function(data){
            console.group(data)
        });

        //Notify the rooms of a new connection
        socket.on('notify_online', function (data) {
            console.log("User -> [ "+data+" ] has connected");
        });

        //Notify the rooms of a disconnect
        socket.on('notify_offline', function (data) {
            console.log("User -> [ "+data+" ] has disconnected");
        });

        socket.on('error', function (data) {
            alert("User -> [ "+data.Name+" ]\n Description: [ "+data.Description+" ]");
            app_chat.isUserLogged=false;
        });
    })
})();

/***************
 *** -> FUNCTIONS
 ****************/

//To properly format the users/clients of the chat
function mapUsers(item){
    return {username: item.Name, status: item.Status};
}

//To properly format the chat messages between users on retrieving
function mapMessages(item){
    console.dir(item);
    return {sender: item.Sender, receiver: item.Receiver, text: item.Text, timestamp: item.Timestamp}
}

//To auto scroll the chat log
function autoScrollChatLog(){
    var chat_log = document.getElementById('chat-log');
    chat_log.scrollTop = chat_log.scrollHeight;
}