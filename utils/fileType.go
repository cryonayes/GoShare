package utils

import (
	"github.com/h2non/filetype"
	"io/ioutil"
	"mime/multipart"
)

var AllowedTypes = map[string]string{
	"css":  "text/css; charset=utf-8",
	"gif":  "image/gif",
	"htm":  "text/html; charset=utf-8",
	"html": "text/html; charset=utf-8",
	"jpeg": "image/jpeg",
	"jpg":  "image/jpeg",
	"js":   "text/plain; charset=utf-8", // To prevent js injection attacks
	"json": "application/json",
	"pdf":  "application/pdf",
	"png":  "image/png",
	"svg":  "image/svg+xml",
	"webp": "image/webp",
	"xml":  "text/xml; charset=utf-8",
	"3gp":  "video/3gpp",
	"mp4":  "video/mp4",
	"webm": "video/webm",
}

func CheckFileType(file *multipart.FileHeader) (string, error) {
	mFile, err := file.Open()
	if err != nil {
		return "", NewError(InvalidFileType)
	}
	buf, _ := ioutil.ReadAll(mFile)

	ftype, err := filetype.Get(buf)
	if err != nil {
		return "", NewError(InvalidFileType)
	}

	if AllowedTypes[ftype.Extension] == "" {
		return "", NewError(InvalidFileType)
	}
	return ftype.Extension, nil
}
