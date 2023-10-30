package common

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

type HTTPResponse struct {
	Success bool        `json:"success"`
	Content interface{} `json:"content"`
	Error   string      `json:"error"`
}

func Wrap(trace string, err error) error {
	return fmt.Errorf("%s -> %w", trace, err)
}

func Hash(data string, salt string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data + salt))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
