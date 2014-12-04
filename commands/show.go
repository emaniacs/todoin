package commands

import (
	"fmt"
	"github.com/emaniacs/todoin/db"
	"os"
	"strconv"
)

var Key int

func init() {
	if len(os.Args) > 2 {
		Key, _ = strconv.Atoi(os.Args[2])
	} else {
		Key = 0
	}
}

func Read() (int, string) {
	if Key < 1 {
		return -1, fmt.Sprintf("Invalid key %d", Key)
	}
	task := db.ByKey(Key)

	// status := (map[bool]string{true: " [done]", false: ""})[task.Status == 1]
	status := ""
	if task.Status == 1 {
		status = " [done]"
	}

	msg := fmt.Sprintf("%s%s", task.Value, status)

	return 0, msg
}
