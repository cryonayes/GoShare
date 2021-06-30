package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

func GetMD5String(str string) string {
	h := md5.New()
	_, _ = io.WriteString(h, str)
	hashed := fmt.Sprintf("%x", h.Sum(nil))
	return hashed
}
