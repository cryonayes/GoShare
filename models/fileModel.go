package models

type FileModel struct {
	FileName    string `json:"file_name"`
	ContentType string `json:"content_type"`
	FileSize    uint32 `json:"file_size"`
}

const (
	ConstFilename    = "file_name"
	ConstContentType = "content_type"
	ConstFileSize    = "file_size"
)
