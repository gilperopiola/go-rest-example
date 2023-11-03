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

const (
	Reset       = "\033[0m"
	Red         = "\033[31m"
	Green       = "\033[32m"
	Yellow      = "\033[33m"
	Blue        = "\033[34m"
	Magenta     = "\033[35m"
	Cyan        = "\033[36m"
	White       = "\033[37m"
	BlueBold    = "\033[34;1m"
	MagentaBold = "\033[35;1m"
	RedBold     = "\033[31;1m"
	YellowBold  = "\033[33;1m"
	WhiteBold   = "\033[37;1m"
	GreenBold   = "\033[32;1m"
	Gray        = "\033[90m"
)
