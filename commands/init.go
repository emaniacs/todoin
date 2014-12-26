package commands

import (
	"github.com/emaniacs/todoin/db"
)

func init() {
	Register("init", &Command{
		Usage: func() string {
			return "Usage of init"
		},
		Run: func(args []string) int {
			db.TableRemove("task")
			db.TableCreate("task")
			return 0
		},
	})
}
