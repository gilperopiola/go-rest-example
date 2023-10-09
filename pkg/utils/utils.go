package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
)

// General utils

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

// Environment utils

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func GetEnvBool(key string, fallback bool) bool {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value == "true" || value == "1"
}

func GetEnvInt(key string, fallback int) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return intValue
}
