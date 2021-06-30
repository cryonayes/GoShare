package errorUtil

import (
	"encoding/json"
	"errors"
)

type ApiError struct {
	Error string `json:"error"`
}

const (
	UserNotFound        = "User not found!"
	InvalidCredentials  = "Invalid credentials!"
	LoginFailed         = "Login failed!"
	RegisterFailed      = "Registeration failed!"
	UserAlreadyExists   = "User already exists!"
	Unauthenticated     = "Unauthenticated!"
	Unauthorized        = "Unauthorized!"
	UploadError         = "Error while uploading file!"
	FileSavingError     = "Error while saving data!"
	InvalidUsername     = "Username is invalid!"
	InvalidPassword     = "Password is invalid!"
	DatabaseConnErr     = "Database connection failed!"
	RequestError        = "Invalid request!"
	InternalServerError = "Internal server error!"
)

func (e *ApiError) String() string {
	out, _ := json.Marshal(e)
	return string(out)
}

func NewError(error string) error {
	return errors.New(error)
}

func NewJSONError(error string) *ApiError {
	err := &ApiError{Error: error}
	return err
}