package commands

import (
	"fmt"
	"github.com/emaniacs/todoin/db"
	"github.com/emaniacs/todoin/utils"
	"os"
)

func init() {
	Register("set", &Command{
		Usage: func() string {
			return "Usage of set"
		},
		Run: func(args []string) int {
			if len(args) < 1 {
				fmt.Fprintf(os.Stderr, "Usage: %s key <options>\n", appName)
				return 255
			}
			key, err := utils.IsNumeric(args[0])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Invalid key")
				return 255
			}
			if !db.Exist(key) {
				fmt.Fprintln(os.Stderr, "Task not found.")
				return 255
			}
			argsFlag := parseFlag("get")
			argsFlag.Flag.Parse(args[1:])

			task := db.ByKey(key)[0]
			db.SyncTask(task, argsFlag.Task)

			if db.Update(key, task) != nil {
				fmt.Fprintln(os.Stdout, "Fail")
				return 255
			}

			fmt.Fprintln(os.Stdout, "Success")
			return 0
		},
	})
}
