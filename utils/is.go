package utils

import (
	"strconv"
	"strings"
)

func IsNumeric(val string) (int, error) {
	return strconv.Atoi(val)
}

func IsDone(val string) (int, bool) {
	ok := false
	status := -1

	if strings.HasPrefix(val, "!") {
		val = val[1:]
	}

	if val == "ok" || val == "o" {
		status = 1
		ok = true
	} else if val == "ko" || val == "k" {
		status = 0
		ok = true
	}
	return status, ok
}

func IsAssignBy(arg string) bool {
	return strings.HasPrefix(arg, "@")
}

func IsAssignTo(arg string) bool {
	return strings.HasPrefix(arg, "$")
}

func IsDueDate(arg string) bool {
	return strings.HasPrefix(arg, "?")
}
