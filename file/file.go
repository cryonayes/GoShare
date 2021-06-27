package file

type fileStruct struct {
	FileName    string `json:"file_name"`
	ContentType string `json:"content_type"`
	FileSize    uint32 `json:"file_size"`
}

const (
	fileName    = "file_name"
	contentType = "content_type"
	fileSize    = "file_size"
)
