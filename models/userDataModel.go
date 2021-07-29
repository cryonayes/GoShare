package app_models

type UserDataModel struct {
	Name string `json:"name"`
	Lastname string `json:"lastname"`
	Files []UserFileModel `json:"files"`
}