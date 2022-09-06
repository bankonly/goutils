package utils

// This function should not be used in any context
// Must use in helper function or validator without context
// but some case you can use this function

import (
	"fmt"
	"strconv"
)

type P struct {
	Message   string
	ErrorCode string
}

func Panic(statusCode int, p P) {
	panic(fmt.Sprintf("%s-%s::%s", strconv.Itoa(statusCode), p.ErrorCode, p.Message))
}
