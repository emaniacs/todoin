package commands

import (
	"fmt"
	"github.com/emaniacs/todoin/db"
	"os"
	"strconv"
)

var Key int

func init() {
	Key = 1
	if len(os.Args) < 3 {
		Key = 0
	}
}

func Show() (int, []string) {
	msg := []string{}
	if Key < 1 {
		msg = append(msg, "Not enough argument.")
		return -1, msg
	}
	var tasks []*db.Task

	for {
		if isDone(os.Args[2]) {
			tasks = db.ByStatus(Key)
		} else if isNumeric(os.Args[2]) {
			tasks = db.ByKey(Key)
		}
		break
	}

	for key := range tasks {
		task := tasks[key]
		status := ""
		if task.Status == 1 {
			status = " [done]"
		}

		// TODO: format output
		msg = append(msg, fmt.Sprintf("%s%s", task.Value, status))
	}

	return 0, msg
}

func isNumeric(val string) bool {
	key, err := strconv.Atoi(val)
	Key = key
	if err == nil {
		return true
	}
	return false
}

func isDone(val string) bool {
	if val == "done" || val == "d" {
		Key = 1
		return true
	} else if val == "undone" || val == "u" {
		Key = 0
		return true
	}
	return false
}
