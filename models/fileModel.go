package app_models

import "time"

type FileModel struct {
	OrigFileName   string    `json:"orig_file_name"`
	HashedFileName string    `json:"hashed_file_name" gorm:"unique"`
	AccessCode     string    `json:"access_code"`
	FileType       string    `json:"file_type"`
	FileSize       int64     `json:"file_size"`
	Owner          string    `json:"owner"`
	IsEncrypted    bool      `json:"is_encrypted"`
	CreationDate   time.Time `json:"creation_date"`
}
