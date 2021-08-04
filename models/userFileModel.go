package app_models

import "time"

type UserFileModel struct {
	OrigFileName string    `json:"filename"`
	FileType     string    `json:"file_type"`
	FileSize     int64     `json:"file_size"`
	AccessCode   string    `json:"access_code"`
	Owner        string    `json:"owner"`
	CreationDate time.Time `json:"creation_date"`
}
