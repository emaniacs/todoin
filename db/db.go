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

	_, err = conn.Exec("CREATE TABLE IF NOT EXISTS task(id INTEGER PRIMARY KEY, value TEXT, status INT)")
	checkError(err)
}

func ByKey(key int) *Task {
	conn, err := createconnection()
	checkError(err)

	task := new(Task)
	sql := "select id, value, status from task WHERE id = ?"
	stmt, err := conn.Prepare(sql)
	checkError(err)
	defer stmt.Close()

	err = stmt.QueryRow(key).Scan(&task.Id, &task.Value, &task.Status)
	// TODO: check for empty row
	return task
}

func Insert(task *Task) (int, string) {
	conn, err := createconnection()
	checkError(err)

	sql := "insert into task (value, status) values (?, ?)"
	stmt, err := conn.Prepare(sql)
	checkError(err)
	defer stmt.Close()

	// TODO: check duplicate
	res, err := stmt.Exec(task.Value, task.Status)
	if err != nil {
		return -1, "Failed while insert"
	}

	task.Id, _ = res.LastInsertId()
	return 1, fmt.Sprintf("Key is %d", task.Id)
}
