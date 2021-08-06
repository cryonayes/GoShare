package app_models

import "encoding/json"

type User struct {
	Id       uint   `json:"id"`
	Email    string `json:"email" gorm:"unique"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Password string `json:"password"`
}

type UserRegister struct {
	Name           string `json:"name"`
	LastName       string `json:"lastname"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	PasswordRepeat string `json:"passwordRepeat"`
}

func (d *User) String() string {
	marshal, err := json.Marshal(d)
	if err != nil {
		return "Error while converting to string"
	}
	return string(marshal)
}
