package db

func TableCreate(name string) {
	conn, err := createconnection()
	checkError(err)

	_, err = conn.Exec("CREATE TABLE IF NOT EXISTS " + name + "(id INTEGER PRIMARY KEY, value TEXT, status INT, assignby STRING, assignto STRING, duedate STRING, filename STRING, line INTEGER)")
	checkError(err)
}

func TableRemove(name string) {
	conn, err := createconnection()
	checkError(err)

	_, err = conn.Exec("DROP TABLE IF EXISTS " + name)
	checkError(err)
}

func TableExist(name string) bool {
	conn, err := createconnection()
	checkError(err)

	sql := "SELECT name FROM  sqlite_master  WHERE type = 'table' AND name = ?"
	stmt, err := conn.Prepare(sql)
	checkError(err)
	defer stmt.Close()

	var tbl string
	stmt.QueryRow(name).Scan(&tbl)

	if tbl == name {
		return true
	}

	return false
}
