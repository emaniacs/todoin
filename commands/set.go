package commands

import (
	"github.com/emaniacs/todoin/db"
	"os"
	"strings"
)

// TODO: set os.Args as arguments
func parseSetArgs(task *db.Task) {
	for _, arg := range os.Args[3:] {
		if strings.HasPrefix(arg, "@") {
			task.AssignBy = arg[1:]
		} else if strings.HasPrefix(arg, "$") {
			task.AssignTo = arg[1:]
		} else if strings.HasPrefix(arg, "?") {
			task.DueDate = arg[1:]
		} else if strings.HasPrefix(arg, "!") {
			if arg[1:] == "ok" || arg[1:] == "1" {
				task.Status = 1
			} else {
				task.Status = 0
			}
		} else if task.Value == "" {
			task.Value = arg
		}
	}
}

// value status
func Set() (int, string) {
	if len(os.Args) < 4 {
		return -1, "Not enough arguments"
	}
	key, err := isNumeric(os.Args[2])
	if err != nil {
		return -1, "Invalid key."
	}

	if !db.Exist(key) {
		return -1, "Task not found."
	}
	task := db.ByKey(key)[0]

	parseSetArgs(task)

	err = db.Update(key, task)

	if err != nil {
		return 1, "Success"
	}

	return -1, "Fail"
}
