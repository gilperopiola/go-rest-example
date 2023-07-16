package utils

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Hash(salt string, data string) string {
	hasher := sha1.New()
	hasher.Write([]byte(salt + data))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func JoinErrors(err1, err2 error) error {
	return fmt.Errorf("%w: %w", err1, err2)
}

func GetIntFromContext(c *gin.Context, key string) (int, error) {

	// Get from context
	value, ok := c.Get(key)
	if value == nil || !ok {
		return 0, fmt.Errorf("error getting %s from context", key)
	}

	// Cast to string
	valueStr, ok := value.(string)
	if !ok {
		return 0, fmt.Errorf("error casting %s from context to string", key)
	}

	// Convert to int
	valueInt, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, fmt.Errorf("error converting %s from string to int", key)
	}

	return valueInt, nil
}

func GetIntFromURLParams(c *gin.Context, key string) (int, error) {

	// Get from params
	value, ok := c.Params.Get(key)
	if !ok {
		return 0, fmt.Errorf("error getting %s from URL params", key)
	}

	// Convert to int
	valueInt, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("error converting %s from string to int", key)
	}

	return valueInt, nil
}
