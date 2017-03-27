package data

import "database/sql"

func DbHandle() *sql.DB {
	db, err := sql.Open("mysql", "goApiUser:password@/goApi")
	checkErr(err)

	err = db.Ping()
	checkErr(err)

	return db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
