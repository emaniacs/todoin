package commands

import (
	"github.com/emaniacs/todoin/db"
	"github.com/emaniacs/todoin/utils"
	"os"
)

// TODO: set os.Args as arguments
func parseAddArgs(task *db.Task) {
	for _, arg := range os.Args[2:] {
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
func Add() (int, string) {
	if len(os.Args) < 3 {
		return -1, "Not enough arguments"
	}

	task := new(db.Task)
	parseAddArgs(task)

	key, msg := db.Insert(task)

	if key >= 1 {
		return 0, msg
	}

	return -1, msg
}
