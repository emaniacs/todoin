package commands

import (
	"fmt"
	"github.com/emaniacs/todoin/db"
	"os"
	"strings"
)

// TODO: set os.Args as arguments
func parseAddArgs() (bool, *db.Task) {
	task := new(db.Task)
	err := false
	sliceStart := 0

	if len(os.Args) >= 4 {
		task.Value = os.Args[2]
		sliceStart = 3
		switch os.Args[3] {
		case "ok", "1":
			task.Status = 1
		case "0":
			task.Status = 0
		default:
			task.Status = 0
		}
	} else if len(os.Args) >= 3 {
		sliceStart = 2
		task.Value = os.Args[2]
		task.Status = 0
	} else {
		err = true
	}

	for _, arg := range os.Args[sliceStart:] {
		if strings.HasPrefix(arg, "@") {
			task.AssignBy = arg[1:]
		} else if strings.HasPrefix(arg, "$") {
			task.AssignTo = arg[1:]
		} else if strings.HasPrefix(arg, "?") {
			task.DueDate = arg[1:]
		}
	}

	return err, task
}

// value status
func Add() (int, string) {
	err, task := parseAddArgs()
	if err {
		return -1, fmt.Sprintf("Not enough arguments", task)
	}

	key, msg := db.Insert(task)

	if key >= 1 {
		return 0, msg
	}

	return -1, msg
}
