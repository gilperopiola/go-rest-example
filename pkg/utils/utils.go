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
	return fmt.Errorf("%w:%w", err1, err2)
}

func GetIntFromContext(c *gin.Context, key string) (int, error) {
	value, ok := c.Get(key)
	if value == nil || !ok {
		return 0, fmt.Errorf("error getting %s from context", key)
	}

	valueStr, ok := value.(string)
	if !ok {
		return 0, fmt.Errorf("error casting %s from context to string", key)
	}

	return strconv.Atoi(valueStr)
}

func GetIntFromURLParams(c *gin.Context, key string) (int, error) {
	value, ok := c.Params.Get(key)
	if !ok {
		return 0, fmt.Errorf("error getting %s from URL params", key)
	}

	valueInt, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("error casting %s from URL params to int", key)
	}

	return valueInt, nil
}
