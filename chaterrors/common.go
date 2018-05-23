package chaterrors

import (
	"log"
	"reflect"
)

// SocketError structure that represents an error in the socket.
type SocketError struct {
	Name        string
	Description string
}

/**
Public Functions
*/

//CheckError function that checks for an error and logs what error occurred.
func CheckError(error error) bool {
	if error != nil {
		log.Printf("%v", error)
		return true
	}
	return false
}

// NewError function that creates a new instance of a SocketError and returns it.
func NewError(name, description string) SocketError {
	error := SocketError{name, description}
	return error
}

// CheckIncorrectString function that checks the value of the string moved.
// If it is empty or is nil, returns true.
// If it is not of the type string, returns true.
// Else it returns false and is correct.
func CheckIncorrectString(data interface{}) bool {
	if data == nil {
		return true
	}
	if reflect.TypeOf(data).String() != "string" {
		return true
	}
	return false
}
