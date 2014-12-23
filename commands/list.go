package commands

import (
	"fmt"
)

func init() {
	Register("list", &Command{
		Usage: func() string {
			return "List of command"
		},
		Run: func(args []string) int {
			for name, _ := range containers {
				fmt.Println(name)
			}

			return 0
		},
	})
}
