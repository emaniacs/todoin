package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func createconnection() (*sql.DB, error) {
	return sql.Open("sqlite3", "tasks.db")
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
	sql := "select id, value, status, assignby, assignto, duedate from task WHERE id = ?"
	stmt, err := conn.Prepare(sql)
	checkError(err)
	defer stmt.Close()

	err = stmt.QueryRow(key).Scan(&task.Id, &task.Value, &task.Status, &task.AssignBy, &task.AssignTo, &task.DueDate)
	// TODO: check for empty row

	var tasks []*Task
	tasks = append(tasks, task)
	return tasks
}

func Insert(task *Task) (int, string) {
	conn, err := createconnection()
	checkError(err)

	sql := "insert into task (value, status, assignby, assignto, duedate) values (?, ?, ?, ?, ?)"
	stmt, err := conn.Prepare(sql)
	checkError(err)
	defer stmt.Close()

	// TODO: check duplicate
	res, err := stmt.Exec(task.Value, task.Status, task.AssignBy, task.AssignTo, task.DueDate)
	if err != nil {
		return -1, "Failed while insert"
	}

	task.Id, _ = res.LastInsertId()
	return 1, fmt.Sprintf("Key is %d", task.Id)
}

func ByStatus(status int) []*Task {
	conn, err := createconnection()
	checkError(err)

	sql := fmt.Sprintf("select id, value, status, assignby, assignto, duedate from task where status = %d", status)
	rows, err := conn.Query(sql)
	checkError(err)
	defer rows.Close()

	var tasks []*Task

	for rows.Next() {
		task := new(Task)
		rows.Scan(&task.Id, &task.Value, &task.Status)
		tasks = append(tasks, task)
	}

	return tasks
}

func GetAll() []*Task {
	conn, err := createconnection()
	checkError(err)

	sql := fmt.Sprintf("select id, value, status, assignby, assignto, duedate  from task")
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

	sql := "select id from task where id = ?"
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

func Update(key int, task *Task) bool {
	conn, err := createconnection()
	checkError(err)

	sql := fmt.Sprintf("UPDATE task SET value = ?, status = ?, assignby = ?, assignto = ?, duedate = ? WHERE id = %d", key)

	stmt, err := conn.Prepare(sql)
	checkError(err)
	defer stmt.Close()

	_, err = stmt.Exec(task.Value, task.Status, task.AssignBy, task.AssignTo, task.DueDate)
	if err != nil {
		return true
	}

	return false
}
