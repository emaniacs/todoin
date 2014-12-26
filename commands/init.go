package commands

import (
	"fmt"
	"github.com/emaniacs/todoin/db"
)

func init() {
	Register("init", &Command{
		Usage: func() string {
			return fmt.Sprintf(`
Initialize the table.
This command will re-create the table (task)
Usage:
	%s init
`, appName)
		},
		Run: func(args []string) int {
			db.TableRemove("task")
			db.TableCreate("task")
			return 0
		},
	})
}
