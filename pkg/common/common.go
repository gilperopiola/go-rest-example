package common

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/gilperopiola/go-rest-example/pkg/common/config"
)

/*----------------------------------------------------------
// Yes, we have a global config and logger. Deal with it.
/-------------------------------------------------------*/

var (
	Cfg    *config.Config
	Logger LoggerI
)

func SetConfig(c *config.Config) {
	Cfg = c
}

func SetLogger(l LoggerI) {
	Logger = l
}

// HTTPResponse is the standard response format for all endpoints.
type HTTPResponse struct {
	Success bool        `json:"success"`
	Content interface{} `json:"content"`
	Error   string      `json:"error"`
}

// Wrap is just a wrapper for fmt.Errorf
func Wrap(trace string, err error) error {
	return fmt.Errorf("%s -> %w", trace, err)
}

// Hash returns a base64 encoded sha256 hash of the data + salt
func Hash(data string, salt string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data + salt))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
