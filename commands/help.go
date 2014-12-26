package commands

import (
	"fmt"
	"os"
)

func init() {
	// register for help
	Register("help", &Command{
		Usage: func() string {
			return fmt.Sprintf("Usage: %s help <command>", appName)
		},
		Run: func(args []string) int {
			if len(args) > 0 {
				cmd, ok := Exist(args[0])
				if ok {
					fmt.Fprintln(os.Stdout, cmd.Usage())
					return 0
				}
				return 255
			}
			fmt.Fprintf(os.Stdout, "Usage: %s help <command>.\n", appName)
			return 0
		},
	})
}
