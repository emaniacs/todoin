package db

import (
	"database/sql"
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

	_, err = conn.Exec("CREATE TABLE IF NOT EXISTS task(id INTEGER PRIMARY KEY, name TEXT, value TEXT, status INT)")
	checkError(err)
}

func ByKey(key int) *Task {
	conn, err := createconnection()
	checkError(err)

	task := new(Task)
	sql := "select id, name, value, status from task WHERE id = ?"
	stmt, err := conn.Prepare(sql)
	checkError(err)
	defer stmt.Close()

	err = stmt.QueryRow(key).Scan(&task.Id, &task.Name, &task.Value, &task.Status)
	// TODO: check for empty row
	return task
}
