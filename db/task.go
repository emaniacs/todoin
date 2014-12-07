package db

type Task struct {
	Id     int64
	Value  string // ""
	Status int    // !

	AssignBy string // @<
	AssignTo string // @>
	DueDate  string // ?
}
