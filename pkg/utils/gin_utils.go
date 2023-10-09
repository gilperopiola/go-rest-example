package utils

import (
	"fmt"
	"strconv"
)

// We introduce these two instead of just using gin.Context to remove the dependency on gin
type ContextGetter interface {
	Get(key string) (interface{}, bool)
}

type ParamsGetter interface {
	Get(name string) (string, bool)
}

func GetIntFromContext(c ContextGetter, key string) (int, error) {

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

func GetIntFromContextURLParams(params ParamsGetter, key string) (int, error) {

	// Get from params
	value, ok := params.Get(key)
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
