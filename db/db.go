package db

import (
	"database/sql"
	"fmt"
	"github.com/emaniacs/todoin/utils"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func createconnection() (*sql.DB, error) {
	dbName := utils.Config.Database.Name
	if dbName == "" {
		dbName = "tasks.db"
	}
	return sql.Open("sqlite3", dbName)
}

func checkError(err error) error {
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func init() {
	conn, err := createconnection()
	checkError(err)

	_, err = conn.Exec("CREATE TABLE IF NOT EXISTS task(id INTEGER PRIMARY KEY, value TEXT, status INT, assignby STRING, assignto STRING, duedate STRING)")
	checkError(err)
}

func ByKey(key int) []*Task {
	conn, err := createconnection()
	checkError(err)

	task := new(Task)
	sql := "SELECT id, value, status, assignby, assignto, duedate FROM task WHERE id = ?"
	stmt, err := conn.Prepare(sql)
	checkError(err)
	defer stmt.Close()

	err = stmt.QueryRow(key).Scan(&task.Id, &task.Value, &task.Status, &task.AssignBy, &task.AssignTo, &task.DueDate)
	// TODO: check for empty row

	var tasks []*Task
	tasks = append(tasks, task)
	return tasks
}

func Insert(task *Task) int64 {
	conn, err := createconnection()
	checkError(err)

	sql := "INSERT INTO task (value, status, assignby, assignto, duedate) VALUES (?, ?, ?, ?, ?)"
	stmt, err := conn.Prepare(sql)
	checkError(err)
	defer stmt.Close()

	// TODO: check duplicate
	res, err := stmt.Exec(task.Value, task.Status, task.AssignBy, task.AssignTo, task.DueDate)
	if err != nil {
		return -1
	}

	task.Id, _ = res.LastInsertId()
	return task.Id
}

func ByStatus(status int) []*Task {
	conn, err := createconnection()
	checkError(err)

	sql := fmt.Sprintf("SELECT id, value, status, assignby, assignto, duedate FROM task WHERE status = %d", status)
	rows, err := conn.Query(sql)
	checkError(err)
	defer rows.Close()

	var tasks []*Task

	for rows.Next() {
		task := new(Task)
		rows.Scan(&task.Id, &task.Value, &task.Status, &task.AssignBy, &task.AssignTo, &task.DueDate)
		tasks = append(tasks, task)
	}

	return tasks
}

func GetAll() []*Task {
	conn, err := createconnection()
	checkError(err)

	sql := fmt.Sprintf("SELECT id, value, status, assignby, assignto, duedate  FROM task")
	rows, err := conn.Query(sql)
	checkError(err)
	defer rows.Close()

	var tasks []*Task

	for rows.Next() {
		task := new(Task)
		rows.Scan(&task.Id, &task.Value, &task.Status, &task.AssignBy, &task.AssignTo, &task.DueDate)
		tasks = append(tasks, task)
	}

	return tasks
}

func Exist(key int) bool {
	conn, err := createconnection()
	checkError(err)

	sql := "SELECT id FROM task WHERE id = ?"
	stmt, err := conn.Prepare(sql)
	checkError(err)
	defer stmt.Close()

	var id int
	stmt.QueryRow(key).Scan(&id)

	if id == key {
		return true
	}

	return false
}

func Update(key int, task *Task) error {
	conn, err := createconnection()
	checkError(err)

	sql := fmt.Sprintf("UPDATE task SET value = ?, status = ?, assignby = ?, assignto = ?, duedate = ? WHERE id = %d", key)

	stmt, err := conn.Prepare(sql)
	checkError(err)
	defer stmt.Close()

	_, err = stmt.Exec(task.Value, task.Status, task.AssignBy, task.AssignTo, task.DueDate)

	return err
}

func ByUser(user, assign string) []*Task {
	conn, err := createconnection()
	checkError(err)

	var tasks []*Task
	var where string
	if assign == "@" {
		where = "assignby = ?"
	} else if assign == "$" {
		where = "assignto = ?"
	} else {
		return tasks
	}

	sql := "SELECT id, value, status, assignby, assignto, duedate FROM task WHERE " + where
	stmt, err := conn.Prepare(sql)
	checkError(err)
	defer stmt.Close()

	rows, err := stmt.Query(user)
	checkError(err)
	defer rows.Close()

	for rows.Next() {
		task := new(Task)
		rows.Scan(&task.Id, &task.Value, &task.Status, &task.AssignBy, &task.AssignTo, &task.DueDate)
		tasks = append(tasks, task)
	}

	return tasks
}

func GetWhere(user, assign string) []*Task {
	conn, err := createconnection()
	checkError(err)

	var tasks []*Task
	var where string
	if assign == "@" {
		where = "assignby = ?"
	} else if assign == "$" {
		where = "assignto = ?"
	} else {
		return tasks
	}

	sql := "SELECT id, value, status, assignby, assignto, duedate FROM task WHERE " + where
	stmt, err := conn.Prepare(sql)
	checkError(err)
	defer stmt.Close()

	rows, err := stmt.Query(user)
	checkError(err)
	defer rows.Close()

	for rows.Next() {
		task := new(Task)
		rows.Scan(&task.Id, &task.Value, &task.Status, &task.AssignBy, &task.AssignTo, &task.DueDate)
		tasks = append(tasks, task)
	}

	return tasks
}

func GetWheres(where string) []*Task {
	conn, err := createconnection()
	checkError(err)

	sql := fmt.Sprintf("SELECT id, value, status, assignby, assignto, duedate FROM task WHERE  %s", where)
	rows, err := conn.Query(sql)
	checkError(err)
	defer rows.Close()

	var tasks []*Task

	for rows.Next() {
		task := new(Task)
		rows.Scan(&task.Id, &task.Value, &task.Status, &task.AssignBy, &task.AssignTo, &task.DueDate)
		tasks = append(tasks, task)
	}

	return tasks
}

func DeleteById(key int64) bool {
	conn, err := createconnection()
	checkError(err)

	sql := "DELETE FROM task WHERE id = ?"
	stmt, err := conn.Prepare(sql)
	checkError(err)
	defer stmt.Close()

	// TODO: check duplicate
	_, err = stmt.Exec(key)
	if err != nil {
		return false
	}
	return true
}
