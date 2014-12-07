package commands

import (
	"fmt"
	"github.com/emaniacs/todoin/db"
	"os"
	"strconv"
)

func Show() (int, []string) {
	msg := []string{}
	var tasks []*db.Task

	for {
		length := len(os.Args)
		if length == 2 {
			tasks = db.GetAll()
		} else if arg, ok := isDone(os.Args[2]); ok == true {
			tasks = db.ByStatus(arg)
		} else if arg, ok := isNumeric(os.Args[2]); ok == true {
			tasks = db.ByKey(arg)
		} else {
			fmt.Println("Uknown command \"" + os.Args[2] + "\"")
		}
		break
	}

	for key := range tasks {
		task := tasks[key]
		status := "not done"
		if task.Status == 1 {
			status = "done"
		}

		// TODO: format output
		msg = append(msg, fmt.Sprintf("(%d) \"%s\" [%s]", task.Id, task.Value, status))
	}

	return 0, msg
}

func isNumeric(val string) (int, bool) {
	arg, err := strconv.Atoi(val)
	if err == nil {
		return arg, true
	}
	return arg, false
}

func isDone(val string) (int, bool) {
	ret := false
	arg := -1
	if val == "done" || val == "d" {
		arg = 1
		ret = true
	} else if val == "undone" || val == "u" {
		arg = 0
		ret = true
	}
	return arg, ret
}
