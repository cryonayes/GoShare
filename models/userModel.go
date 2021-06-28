package models

import "encoding/json"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const (
	Username = "username"
	Password = "password"
)

func (d *User) String() string {
	marshal, err := json.Marshal(d)
	if err != nil {
		return "Error while converting to string"
	}
	return string(marshal)
}
