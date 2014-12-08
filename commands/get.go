package commands

import (
	"fmt"
	"github.com/emaniacs/todoin/db"
	"github.com/emaniacs/todoin/utils"
	"os"
)

func Get() (int, []string) {
	msg := []string{}
	var tasks []*db.Task

	for {
		length := len(os.Args)
		if length == 2 {
			tasks = db.GetAll()
		} else if arg, err := utils.IsDone(os.Args[2]); err == true {
			tasks = db.ByStatus(arg)
		} else if arg, err := utils.IsNumeric(os.Args[2]); err == nil {
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
