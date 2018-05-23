package chaterrors

import "fmt"

/**
** ErrTextIsEmpty
**/

// ErrTextIsEmpty for errors where the text in a message is empty.
type ErrTextIsEmpty struct {
	Sender   string
	Receiver string
	Text     string
}

func (err ErrTextIsEmpty) Error() string {
	return fmt.Sprintf("The message between [%s] and [%s] is empty: [%s]", err.Sender, err.Receiver, err.Text)
}

// IsErrTextIsEmpty function to check if an error is of the type in the name.
func IsErrTextIsEmpty(err error) bool {
	_, ok := err.(ErrTextIsEmpty)
	return ok
}

/**
** ErrSenderNameIsEmpty
**/

// ErrSenderNameIsEmpty for errors where the name of the sender is empty.
type ErrSenderNameIsEmpty struct {
	Username string
}

func (err ErrSenderNameIsEmpty) Error() string {
	return fmt.Sprintf("The Sender name [%s] is empty", err.Username)
}

// IsErrSenderNameIsEmpty function to check if an error is of the type in the name.
func IsErrSenderNameIsEmpty(err error) bool {
	_, ok := err.(ErrSenderNameIsEmpty)
	return ok
}

/**
** ErrSenderDoesNotExist
**/

// ErrSenderDoesNotExist for errors where the sender who sent the message does not exist.
type ErrSenderDoesNotExist struct {
	Username string
}

func (err ErrSenderDoesNotExist) Error() string {
	return fmt.Sprintf("The Sender [%s] does not exist", err.Username)
}

// IsErrSenderDoesNotExist function to check if an error is of the type in the name.
func IsErrSenderDoesNotExist(err error) bool {
	_, ok := err.(ErrSenderDoesNotExist)
	return ok
}

/**
** ErrReceiverNameIsEmpty
**/

// ErrReceiverNameIsEmpty for errors when the name of the receiver is empty.
type ErrReceiverNameIsEmpty struct {
	Username string
}

func (err ErrReceiverNameIsEmpty) Error() string {
	return fmt.Sprintf("The Receiver name [%s] is empty", err.Username)
}

// IsErrReceiverNameIsEmpty function to check if an error is of the type in the name.
func IsErrReceiverNameIsEmpty(err error) bool {
	_, ok := err.(ErrReceiverNameIsEmpty)
	return ok
}

/**
** ErrReceiverDoesNotExist
**/

// ErrReceiverDoesNotExist for errors when the receiver does not exist.
type ErrReceiverDoesNotExist struct {
	Username string
}

func (err ErrReceiverDoesNotExist) Error() string {
	return fmt.Sprintf("The Receiver [%s] does not exist", err.Username)
}

// IsErrReceiverDoesNotExist function to check if an error is of the type in the name.
func IsErrReceiverDoesNotExist(err error) bool {
	_, ok := err.(ErrReceiverDoesNotExist)
	return ok
}

/**
** ErrClientsAreNotFriends
**/

// ErrClientsAreNotFriends for errors where both clients are not friends.
type ErrClientsAreNotFriends struct {
	Sender   string
	Receiver string
}

func (err ErrClientsAreNotFriends) Error() string {
	return fmt.Sprintf("The Sender [%s] and Receiver [%s] are not friends.", err.Sender, err.Receiver)
}

// IsErrClientsAreNotFriends function to check if an error is of the type in the name.
func IsErrClientsAreNotFriends(err error) bool {
	_, ok := err.(ErrClientsAreNotFriends)
	return ok
}
