package commands

import (
	"fmt"
	"github.com/emaniacs/todoin/db"
	"github.com/emaniacs/todoin/utils"
	"os"
)

func init() {
	Register("del", &Command{
		Usage: func() string {
			return "Usage of del"
		},
		Run: func(args []string) int {
			if len(args) < 1 {
				fmt.Fprintln(os.Stderr, "Not enough argument")
				return 255
			}

			// TODO: remove by column
			key, err := utils.IsNumeric(args[0])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Invalid key")
				return 255
			}
			if !db.Exist(key) {
				fmt.Fprintln(os.Stderr, "Task not found.")
				return 255
			}

			if db.DeleteById(int64(key)) {
				fmt.Fprintln(os.Stdout, "Success")
				return 0
			}
			fmt.Fprintln(os.Stdout, "Fail")
			return 255
		},
	})
}
