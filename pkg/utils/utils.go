package utils

import (
	"crypto/sha1"
	"encoding/base64"
	"strconv"
	"strings"
)

func toString(i int) string {
	return strconv.Itoa(i)
}

func toInt(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func beautifyError(err error) string {
	s := err.Error()

	if strings.Contains(s, "Duplicate entry") {
		duplicateField := strings.Split(s, "'")[3]
		return duplicateField + " already in use"
	}

	return s
}

func Hash(salt string, data string) string {
	hasher := sha1.New()
	hasher.Write([]byte(salt + data))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
