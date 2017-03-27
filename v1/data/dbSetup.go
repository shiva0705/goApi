package data

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/shiva0705/goApi/v1/models"
)

func init() {
	buildTables()
	populateVideos()
}

func buildTables() {
	var db = getDBHadle()
	defer db.Close()

	// create person table
	stmtDrop, err := db.Prepare("Drop table if exists person")
	checkErr(err)

	_, err = stmtDrop.Exec()
	checkErr(err)

	stmtCreatePerson, err := db.Prepare("CREATE TABLE person (id int NOT NULL AUTO_INCREMENT, name varchar(250), url varchar(250), like_count int, dislike_count int, PRIMARY KEY (id))AUTO_INCREMENT=1;")
	checkErr(err)

	_, err = stmtCreatePerson.Exec()
	checkErr(err)

}

func populateVideos() {
	var db = getDBHadle()
	defer db.Close()

	addVideo(db, models.Video{Id: 1, Name: "Ultimate Dog Tease", Url: "https://www.youtube.com/watch?v=nGeKSiCQkPw", LikeCount: 0, DislikeCount: 0})
	addVideo(db, models.Video{Id: 2, Name: "Dog Wants Kitty", Url: "https://www.youtube.com/watch?v=kI4yoXyb1_M", LikeCount: 0, DislikeCount: 0})
}
