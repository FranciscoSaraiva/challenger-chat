**OLX Frontend Chat Challenge**
====================
## Socket documentation
The following document shows the information needed to work with the socket and the various calls the server sends back to the browser.

### 1. Structure
#### Important points to take in consideration
1. **The only type of information the socket needs to work with are objects.**
2. **The only type of information sent inside the object are strings, if it's not a string the socket will return an error back to the browser socket.**


The four named parameters passed over are:


|Parameter|Description|
|---|---|
|**username**|represents the name of the user in the application. Must be 3 digits or higher and have no numeric values.|
|**friendname** |represents the name of a friend of the user.
|**status** |represents the status of the user in the application. *("Online", "Away", "Busy")*|
|**text**|represents a message to be sent to a friend.|

So the client on the browser side must send objects with the format below:

```
{
	username: "John Doe",
	friendname: "James Rolfe",
	status: "Online",
	text: "How ya doin' nerd?"
}
```

Not all fields are required, but these are the only four fields needed to communicate with the socket. Like for example:

##### **Login Example**
```
{
	username: "John Doe"
}
```

Only the username is needed to login, so you only need to send the username of the user that wants to login.

---
### 2. Socket calls
These are the various calls you can make to the server to make things work.

#### **Login**

**Call name**

> login_user

##### **Requires**
- ```username``` - The username of the person that wants to logs in the chat application.

##### **Format Example**

```
{
	username: "John Doe"
}
```

##### **Returns**
- Broadcasts to:
 - **room:** ```chat```,
 - **call:**  ```notify_online```
 - **returns:**  ```username (string)``` of the user that has logged in.
- Emits:
 - **call:** ```user_session```
 - **return:** ```userInfo (object)``` with the properties ```User (string)``` and ```Status (string)```

---
---
---

#### **Logout**

**Call name**

> logout_user

##### **Requires**
- ```username``` - The username of the person that is logged in the application.

##### **Format Example**

```
{
	username: "John Doe"
}
```

##### **Returns**
- Broadcasts to room ```chat```, the call ```notify_status``` and returns the ```userInfo``` which is an array of ```strings``` with the user's ```username``` and the ```status``` "Offline".

- Broadcasts to:
 - **room:** ```chat```,
 - **call:**  ```notify_offline```
 - **returns:**  ```username (string)``` of the user that has logged in.

- Broadcasts to:
  - **room:** ```chat```
  - **call:** ```notify_status```
  - **return:** ```userInfo (object)``` with the ```User (string)``` as the username of the user and ```Status (string)``` with the value "Offline".

---
---
---

#### **Add Friend**

**Call name**

> add_friend

##### **Requires**
- ```username``` - The username of the person that is logged in the application.
- ```friendname``` - The name of the friend the user wishes to add.

##### **Format Example**

```
{
	username: "John Doe",
	friendname: "James Rolfe"
}
```

##### **Returns**
- Emits:
 - **call:** ```get_user_friends```
 - **returns:** ```friends (array)``` which is an ```array``` of ```objects``` with the properties ```Name (string)``` and ```Status (string)```

---
---
---

#### **Get Chat Log**

**Call name**

> get_chat_log

##### **Requires**
- ```username``` - The username of the person that is logged in the application.
- ```friendname``` - The name of the friend the user wants its chat history.

##### **Format Example**

```
{
	username: "John Doe",
	friendname: "James Rolfe"
}
```

##### **Returns**
- The socket joins a room for the two users to chat on.
- Emits:
 - **call:** ```get_chat_log```
 - **return:** ```messages (array)``` which is an ```array``` of ```objects``` with the properties ```Sender (string)``` who sent the message, ```Receiver (string)``` who received the message, ```Text (string)``` the text message and ```Timestamp (string)``` the timestamp of the message at which it was first sent.

---
---
---

#### **Message Friend**

**Call name**

> message_friend

##### **Requires**
- ```username``` - The username of the person that is logged in the application.
- ```friend``` - The name of the friend the user wishes to message.
- ```text``` - The text message the user wants to send to the friend.

##### **Format Example**

```
{
	username: "John Doe",
	friendname: "James Rolfe",
	text: "How ya doin' nerd?"
}
```

##### **Returns**
- Broadcasts to:
 - **room:** ```(assigned room name)```
 - **call:** ```get_new_message```
 - **return:** ```message (object)``` with the properties ```Sender (string)``` who sent the message, ```Receiver (string)``` who received the message, ```Text (string)``` the text message and ```Timestamp``` the timestamp of the message at which it was first sent.

---
---
---

#### **Get Friends List**

**Call name**

> get_user_friends

##### **Requires**
- ```username``` - The username of the person that is logged in the application.

##### **Format Example**

```
{
	username: "John Doe"
}
```

##### **Returns**
- Emits:
 - **call:** ```get_user_friends```
 - **return:** ```friends (array)``` which is an array of ```objects``` with the properties ```Name (string)``` and ```Status (string)```

---
---
---

#### **Change Status**

**Call name**

> change_status

##### **Requires**
- ```username``` - The username of the person that is logged in the application.
- ```status``` - The status the user wants to change to.

##### **Format Example**

```
{
	username: "John Doe",
	status: "Busy"
}
```

##### **Returns**
- Broadcasts to:
 - **room:** ```chat```
 - **call:** ```notify_status```
 - **return:** ```userInfo (object)``` with the properties ```User (string)``` of the user that changed status and ```Status (string)``` the status that user changed to.

**Note the only status options are "Online", "Away" and "Busy"**


---

### 3. Errors
In case something goes wrong with the communication with the server, whenever information is badly sent inside the objects, the server will **always** emit the following:

- Emits:
 - **call:** ```error```
 - **return:** ```error (object)``` with the properties  ```Name (string)``` with the name of the error and ```Description (string)``` with a description of the error.