package user

import (
	"encoding/json"
)

type userData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const (
	username = "username"
	password = "password"
)

func (d *userData) String() string {
	marshal, err := json.Marshal(d)
	if err != nil {
		return "Error while converting to string"
	}
	return string(marshal)
}
