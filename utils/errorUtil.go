package utils

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
)

func (e *ErrorUtil) String() string {
	out, _ := json.Marshal(e)
	return string(out)
}

func NewError(error string) *ErrorUtil {
	err := &ErrorUtil{Error: error}
	return err
}
