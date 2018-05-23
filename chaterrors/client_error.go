package chaterrors

import "fmt"

/**
** ErrClientNameIsEmpty
**/

// ErrClientLoginNameIsEmpty for errors where the client name is empty
type ErrClientLoginNameIsEmpty struct {
	Username string
}

func (err ErrClientLoginNameIsEmpty) Error() string {
	return fmt.Sprintf("The Login name is empty: [%s]", err.Username)
}

// IsErrClientLoginNameIsEmpty function to check if an error is of the type in the name.
func IsErrClientLoginNameIsEmpty(err error) bool {
	_, ok := err.(ErrClientLoginNameIsEmpty)
	return ok
}

/**
** ErrClientLoginNameIsShort
**/

// ErrClientLoginNameIsShort for errors where the client name is less than 3 characters
type ErrClientLoginNameIsShort struct {
	Username string
}

func (err ErrClientLoginNameIsShort) Error() string {
	return fmt.Sprintf("The Login name is too short: [%s]", err.Username)
}

// IsErrClientLoginNameIsShort function to check if an error is of the type in the name.
func IsErrClientLoginNameIsShort(err error) bool {
	_, ok := err.(ErrClientLoginNameIsShort)
	return ok
}

/**
** ErrClientLoginNameHasNumeric
**/

// ErrClientLoginNameHasNumeric for errors where the client name has numeric values.
type ErrClientLoginNameHasNumeric struct {
	Username string
}

func (err ErrClientLoginNameHasNumeric) Error() string {
	return fmt.Sprintf("The Login name is too short: [%s]", err.Username)
}

// IsErrClientLoginNameHasNumeric function to check if an error is of the type in the name.
func IsErrClientLoginNameHasNumeric(err error) bool {
	_, ok := err.(ErrClientLoginNameHasNumeric)
	return ok
}

/**
** ErrClientNotFound
**/

// ErrClientNotFound for errors where the client isn't found in the database.
type ErrClientNotFound struct {
	ID       string
	Username string
}

func (err ErrClientNotFound) Error() string {
	return fmt.Sprintf("\n User does not exist [ID: %s] [Username: %s]\n", err.ID, err.Username)
}

// IsErrClientNotFound function to check if an error is of the type in the name.
func IsErrClientNotFound(err error) bool {
	_, ok := err.(ErrClientNotFound)
	return ok
}

/**
** ErrClientInvalidName
**/

// ErrClientInvalidName for errors where the client has an invalid name.
type ErrClientInvalidName struct {
	Username string
}

func (err ErrClientInvalidName) Error() string {
	return fmt.Sprintf("The username [Username: %s] is not valid", err.Username)
}

// IsErrClientInvalidName function to check if an error is of the type in the name.
func IsErrClientInvalidName(err error) bool {
	_, ok := err.(ErrClientInvalidName)
	return ok
}

/**
** ErrClientNameEmpty
**/

// ErrClientNameEmpty for errors where the client name is empty.
type ErrClientNameEmpty struct {
	Username string
}

func (err ErrClientNameEmpty) Error() string {
	return fmt.Sprintf("The username is empty [%s]", err.Username)
}

// IsErrClientNameEmpty function to check if an error is of the type in the name.
func IsErrClientNameEmpty(err error) bool {
	_, ok := err.(ErrClientNameEmpty)
	return ok
}

/**
** ErrInvalidStatus
**/

// ErrInvalidStatus for errors where the status of the client is invalid.
type ErrInvalidStatus struct {
	Status string
}

func (err ErrInvalidStatus) Error() string {
	return fmt.Sprintf("The status option [%s] is invalid", err.Status)
}

// IsErrInvalidStatus function to check if an error is of the type in the name.
func IsErrInvalidStatus(err error) bool {
	_, ok := err.(ErrInvalidStatus)
	return ok
}

/**
** ErrEmptyFriendName
**/

// ErrEmptyFriendName for errors where the name of the friend is empty.
type ErrEmptyFriendName struct {
	Friendname string
}

func (err ErrEmptyFriendName) Error() string {
	return fmt.Sprintf("The friend name is empty [%s]", err.Friendname)
}

// IsErrEmptyFriendName function to check if an error is of the type in the name.
func IsErrEmptyFriendName(err error) bool {
	_, ok := err.(ErrEmptyFriendName)
	return ok
}

/**
** ErrFriendNotFound
**/

// ErrFriendNotFound for errors where the friend is not found in the database.
type ErrFriendNotFound struct {
	Friendname string
}

func (err ErrFriendNotFound) Error() string {
	return fmt.Sprintf("The friend [%s] was not found", err.Friendname)
}

// IsErrFriendNotFound function to check if an error is of the type in the name.
func IsErrFriendNotFound(err error) bool {
	_, ok := err.(ErrFriendNotFound)
	return ok
}

/**
** ErrUsersAreFriends
**/

// ErrUsersAreFriends for errors where both users are already friends.
type ErrUsersAreFriends struct {
	Username   string
	Friendname string
}

func (err ErrUsersAreFriends) Error() string {
	return fmt.Sprintf("The user [%s] and the user [%s] are already friends", err.Username, err.Friendname)
}

// IsErrUsersAreFriends function to check if an error is of the type in the name.
func IsErrUsersAreFriends(err error) bool {
	_, ok := err.(ErrUsersAreFriends)
	return ok
}
