package data

import (
	"database/sql"

	"github.com/shiva0705/goApi/v1/models"

	_ "github.com/go-sql-driver/mysql"
)

func addVideo(db *sql.DB, video models.Video) (err error) {
	tx, err := db.Begin()
	checkErr(err)

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	if _, err = tx.Exec("Insert Into person Values (?, ?, ?, ?, ?)", 0, video.Name, video.Url, video.LikeCount, video.DislikeCount); err != nil {
		return
	}
	return
}

func GetVideos() models.Videos {
	var db = getDBHadle()
	defer db.Close()

	rows, err := db.Query("Select * from person")
	checkErr(err)

	var videos models.Videos
	var video models.Video

	for rows.Next() {
		err = rows.Scan(&video.Id, &video.Name, &video.Url, &video.LikeCount, &video.DislikeCount)
		videos = append(videos, video)
		checkErr(err)
	}

	defer rows.Close()

	return videos
}

func UpdateFeedback(f models.Feedback) {
	var db = getDBHadle()
	defer db.Close()

	var video models.Video

	err := db.QueryRow("Select id, name, url, like_count, dislike_count from person where id=?", f.VideoId).Scan(&video.Id, &video.Name, &video.Url, &video.LikeCount, &video.DislikeCount)
	checkErr(err)

	if f.Like == true {
		video.LikeCount++
	} else {
		video.DislikeCount++
	}

	stmt, err := db.Prepare("Update person Set like_count =?, dislike_count =? where id =?")
	checkErr(err)
	_, err = stmt.Exec(video.LikeCount, video.DislikeCount, video.Id)
	checkErr(err)
}
