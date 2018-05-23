package chaterrors

import "fmt"

// ErrSessionNotFound for errors where the session does not exist.
type ErrSessionNotFound struct {
	Username string
}

func (err ErrSessionNotFound) Error() string {
	return fmt.Sprintf("User [%s] does not have a valid Section", err.Username)
}

// IsErrSessionNotFound function to check if an error is of the type in the name.
func IsErrSessionNotFound(err error) bool {
	_, ok := err.(ErrSessionNotFound)
	return ok
}
