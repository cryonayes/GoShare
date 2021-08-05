package app_models

import "time"

type FileModel struct {
	OrigFileName   string    `json:"orig_file_name"`
	HashedFileName string    `json:"hashed_file_name" gorm:"unique"`
	AccessCode     string    `json:"access_code"`
	Shared         bool      `json:"shared"`
	AccessToken    string    `json:"access_token"`
	ShareTime      time.Time `json:"share_time"`
	FileType       string    `json:"file_type"`
	FileSize       int64     `json:"file_size"`
	Owner          string    `json:"owner"`
	IsEncrypted    bool      `json:"is_encrypted"`
	CreationDate   time.Time `json:"creation_date"`
}

type FileShareDatas struct {
	AccessCode string `json:"accesscode"`
	ShareTime  string `json:"sharetime"`
}

type FileAccessCode struct {
	AccessCode string `json:"accesscode"`
}

type FileAccessLink struct {
	AccessLink string `json:"accesslink"`
}
