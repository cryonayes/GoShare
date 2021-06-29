package errorUtil

import (
	"encoding/json"
)

type ErrorUtil struct {
	Error string `json:"error"`
}

const (
	UserNotFound       = "User not found!"
	InvalidCredentials = "Invalid credentials!"
	LoginFailed        = "Login failed!"
	RegisterFailed     = "Registeration failed!"
	UserAlreadyExists  = "User already exists!"
	Unauthenticated    = "Unauthenticated!"
	Unauthorized       = "Unauthorized!"
)

func (e *ErrorUtil) String() string {
	out, _ := json.Marshal(e)
	return string(out)
}

func NewError(error string) *ErrorUtil {
	err := &ErrorUtil{Error: error}
	return err
}
