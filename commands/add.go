package commands

import (
	"fmt"
	"github.com/emaniacs/todoin/db"
	"os"
)

var task *db.Task
var err bool

func init() {
	task = new(db.Task)
	err = false
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
	} else if len(os.Args) == 2 {
		task.Value = os.Args[2]
		task.Status = 0
	} else {
		err = true
	}
}

// value status
func Add() (int, string) {
	if err {
		return -1, fmt.Sprintf("Not enough arguments")
	}

	key, msg := db.Insert(task)

	if key >= 1 {
		return 0, msg
	}

	return -1, msg
}
