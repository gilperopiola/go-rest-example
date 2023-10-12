package common

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

type QueryOption func(*string)

// Wrap should actaully be called WrapError, but it's too long
func Wrap(err1, err2 error) error {
	return fmt.Errorf("%s: %w", err1.Error(), err2)
}

// Hash hashes
func Hash(salt string, data string) string {
	hasher := sha256.New()
	hasher.Write([]byte(salt + data))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
