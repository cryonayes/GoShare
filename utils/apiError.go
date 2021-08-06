package utils

import (
	"errors"
)

const (
	UserNotFound        = "User not found!"
	InvalidCredentials  = "Invalid credentials!"
	LoginFailed         = "Login failed!"
	PleaseLogin         = "Please login!"
	RegisterFailed      = "Registeration failed!"
	UserAlreadyExists   = "User already exists!"
	Authenticated       = "Authenticated!"
	InvalidFileCode     = "Invalid access code!"
	FileNotExists       = "File doesn't exists!"
	Unauthenticated     = "Unauthenticated!"
	Unauthorized        = "Unauthorized!"
	UploadError         = "Error while uploading file!"
	FileSavingError     = "Error while saving data!"
	FileShared          = "File shared!"
	FileShareExpired    = "Expired link!"
	ErrorWhileDeleting  = "Error while deleting file from server!"
	FileDeleted         = "File deleted!"
	FileUpdated         = "File updated!"
	ErrorWhileUploading = "Error while uploading file!"
	ErrorWhileUnshare   = "File update failed!"
	InvalidTimeValue    = "Invalid time value!"
	InvalidUsername     = "Username is invalid!"
	InvalidEmail        = "Email is invalid!"
	InvalidPassword     = "Password is invalid!"
	InvalidName         = "Name is invalid!"
	InvalidLastname     = "Lastname is invalid!"
	PasswordRepeatWrong = "Passwords doesn't match!"
	UserRegistered      = "User registered!"
	DatabaseConnErr     = "Database connection failed!"
	RequestError        = "Invalid request!"
	InternalServerError = "Internal server error!"
	InvalidFileType     = "File type not allowed!"
)

func NewError(error string) error {
	return errors.New(error)
}
