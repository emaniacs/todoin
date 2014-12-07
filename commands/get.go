package commands

import (
	"fmt"
	"github.com/emaniacs/todoin/db"
	"os"
	"strconv"
)

func Get() (int, []string) {
	msg := []string{}
	var tasks []*db.Task

	for {
		length := len(os.Args)
		if length == 2 {
			tasks = db.GetAll()
		} else if arg, err := isDone(os.Args[2]); err == true {
			tasks = db.ByStatus(arg)
		} else if arg, err := isNumeric(os.Args[2]); err == nil {
			tasks = db.ByKey(arg)
		} else {
			fmt.Println("Uknown command \"" + os.Args[2] + "\"")
		}
		break
	}

	for key := range tasks {
		task := tasks[key]
		status := "!ko"
		if task.Status == 1 {
			status = "!ok"
		}

		assignby := ""
		if task.AssignBy != "" {
			assignby = "@" + task.AssignBy
		}

		assignto := ""
		if task.AssignTo != "" {
			assignto = "$" + task.AssignTo
		}

		duedate := ""
		if task.DueDate != "" {
			duedate = "?" + task.DueDate
		}

		// TODO: format output
		msg = append(msg,
			fmt.Sprintf("(%d) \"%s\" %s %s %s %s", task.Id, task.Value, status, assignby, assignto, duedate))
	}

	return 0, msg
}

func isNumeric(val string) (int, error) {
	return strconv.Atoi(val)
}

func isDone(val string) (int, bool) {
	err := false
	arg := -1
	if val == "ok" || val == "o" {
		arg = 1
		err = true
	} else if val == "ko" || val == "k" {
		arg = 0
		err = true
	}
	return arg, err
}
