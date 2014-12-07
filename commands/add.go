package commands

import (
	"fmt"
	"github.com/emaniacs/todoin/db"
	"os"
)

func parseAddArgs() (bool, *db.Task) {
	task := new(db.Task)
	err := false
	if len(os.Args) == 4 {
		task.Value = os.Args[2]
		switch os.Args[3] {
		case "done", "1":
			task.Status = 1
		case "0":
			task.Status = 0
		default:
			task.Status = 0
		}
	} else if len(os.Args) == 3 {
		task.Value = os.Args[2]
		task.Status = 0
	} else {
		err = true
	}

	return err, task
}

// value status
func Add() (int, string) {
	err, task := parseAddArgs()
	if err {
		return -1, fmt.Sprintf("Not enough arguments")
	}

	key, msg := db.Insert(task)

	if key >= 1 {
		return 0, msg
	}

	return -1, msg
}
