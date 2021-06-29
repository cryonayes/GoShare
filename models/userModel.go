package models

import "encoding/json"

type User struct {
	Id       uint   `json:"id"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

type UserRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (d *User) String() string {
	marshal, err := json.Marshal(d)
	if err != nil {
		return "Error while converting to string"
	}
	return string(marshal)
}
