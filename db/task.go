package db

import (
	"strconv"
)

type Task struct {
	Id     int64
	Value  string // ""
	Status int    // !

	AssignBy string // @<
	AssignTo string // @>
	DueDate  string // ?

	Filename string
	Line     int
}

func SyncTask(task *Task, tasks map[string]*string) {
	for k, v := range tasks {
		if *v == "" {
			continue
		}
		if k == "assignby" {
			task.AssignBy = *v
		} else if k == "assignto" {
			task.AssignTo = *v
		} else if k == "status" {
			status, err := strconv.Atoi(*v)
			if err == nil {
				task.Status = status
			}
		} else if k == "value" {
			task.Value = *v
		} else if k == "duedate" {
			task.DueDate = *v
		}
	}
}
