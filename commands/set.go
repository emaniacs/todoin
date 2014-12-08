package commands

import (
	"github.com/emaniacs/todoin/db"
	"github.com/emaniacs/todoin/utils"
	"os"
)

// TODO: set os.Args as arguments
func parseSetArgs(task *db.Task) {
	for _, arg := range os.Args[3:] {
		if utils.IsAssignBy(arg) {
			task.AssignBy = arg[1:]
		} else if utils.IsAssignTo(arg) {
			task.AssignTo = arg[1:]
		} else if utils.IsDueDate(arg) {
			task.DueDate = arg[1:]
		} else if status, err := utils.IsDone(arg); err != false {
			task.Status = status
		} else if task.Value == "" {
			task.Value = arg
		}
	}
}

// value status
func Set() (int, string) {
	if len(os.Args) < 4 {
		return -1, "Not enough arguments"
	}
	key, err := utils.IsNumeric(os.Args[2])
	if err != nil {
		return -1, "Invalid key."
	}

	if !db.Exist(key) {
		return -1, "Task not found."
	}
	task := db.ByKey(key)[0]
	// temporary
	tmp := task.Value
	task.Value = ""

	parseSetArgs(task)
	if task.Value == "" {
		task.Value = tmp
	}

	err = db.Update(key, task)

	if err != nil {
		return -1, "Fail"
	}

	return 1, "Success"
}
