package main

import (
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
