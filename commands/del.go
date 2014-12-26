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
			return fmt.Sprintf(`Remove task based on key
Usage:
	%s del <key>
Example:
	$ %s del 10 
	# delete multiple task
	$ %s del 10 11 12 13
			`, appName, appName, appName)
		},
		Run: func(args []string) int {
			if len(args) < 1 {
				fmt.Fprintln(os.Stderr, "Not enough argument")
				return 255
			}

			// TODO: remove by column
			for _, val := range args {
				key, err := utils.IsNumeric(val)
				if err != nil {
					fmt.Fprintf(os.Stdout, "(%s) Invalid key\n", val)
					continue
				}
				if !db.Exist(key) {
					fmt.Fprintf(os.Stdout, "(%d) Task not found\n", key)
					continue
				}

				if db.DeleteById(int64(key)) {
					fmt.Fprintf(os.Stdout, "(%d) Success\n", key)
				} else {
					fmt.Fprintf(os.Stdout, "(%d) Fail\n", key)
				}
			}
			return 0
		},
	})
}
