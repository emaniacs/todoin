package commands

import (
	"fmt"
	"github.com/emaniacs/todoin/db"
	"os"
)

func init() {
	Register("add", &Command{
		Usage: func() string {
			return "Usage of add"
		},
		Run: func(args []string) int {
			if len(args) < 1 {
				fmt.Fprintf(os.Stderr, "Use %s add value <options>\n", appName)
				return 255
			}

			task := new(db.Task)
			argsFlag := parseFlag("get")
			argsFlag.Flag.Parse(args[1:])

			db.SyncTask(task, argsFlag.Task)
			task.Value = args[0]

			id := db.Insert(task)
			if id > 0 {
				fmt.Fprintf(os.Stdout, "Success (%d)\n", id)
				return 0
			}

			fmt.Fprintln(os.Stdout, "Fail")
			return 255
		},
	})
}
