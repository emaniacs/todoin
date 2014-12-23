package commands

import (
	"fmt"
	"github.com/emaniacs/todoin/db"
	"github.com/emaniacs/todoin/utils"
	"os"
	"strconv"
	"strings"
)

func init() {
	Register("get", &Command{
		Usage: func() string {
			return "Usage of get"
		},
		Run: func(args []string) int {
			var tasks []*db.Task
			argsFlag := parseFlag("get")

			if len(args) == 0 {
				tasks = db.GetAll()
			} else if arg, err := utils.IsNumeric(args[0]); err == nil {
				tasks = db.ByKey(arg)
			} else {
				argsFlag.Flag.Parse(args)

				var where []string
				for k, v := range argsFlag.Task {
					if *v == "" {
						continue
					}
					if k == "value" {
						where = append(where, fmt.Sprintf("value LIKE '%%%s%%'", *v))
					} else {
						where = append(where, fmt.Sprintf("%s = '%s'", k, *v))
					}
				}

				if len(where) < 1 {
					fmt.Fprintln(os.Stderr, "Invalid argument.")
					return 255
				}

				tasks = db.GetWheres(strings.Join(where, " AND "))
			}

			if *argsFlag.Verbose {
				fmt.Fprintln(os.Stdout, strings.Join(
					[]string{"id", "value", "status", "assignby", "assignto", "duedate"},
					*argsFlag.Separator,
				))
			}

			for key := range tasks {
				task := tasks[key]
				fmt.Fprintln(os.Stdout, strings.Join(
					[]string{strconv.FormatInt(task.Id, 10), task.Value, strconv.Itoa(task.Status), task.AssignBy, task.AssignTo, task.DueDate},
					*argsFlag.Separator,
				))
			}

			return 0
		},
	})
}
