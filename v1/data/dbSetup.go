package data

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	db, err := sql.Open("mysql", "goApiUser:password@/goApi")
	checkErr(err)

	defer db.Close()

	err = db.Ping()
	checkErr(err)

	// create person table
	stmtDrop, err := db.Prepare("Drop table if exists person")
	checkErr(err)

	_, err = stmtDrop.Exec()
	checkErr(err)

	stmtCreatePerson, err := db.Prepare("CREATE TABLE person (id int NOT NULL AUTO_INCREMENT, name varchar(250), url varchar(250), like_count int, dislike_count int, PRIMARY KEY (id));")
	checkErr(err)

	_, err = stmtCreatePerson.Exec()
	checkErr(err)

	fmt.Println("Person table created")

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
