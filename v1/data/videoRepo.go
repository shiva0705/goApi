package data

import (
	"github.com/shiva0705/goApi/v1/models"

	_ "github.com/go-sql-driver/mysql"
)

func addVideo(video models.Video) {
	var db = getDBHadle()
	defer db.Close()

	stmt, err := db.Prepare("Insert Into person Values (?, ?, ?, ?, ?)")
	checkErr(err)

	_, err = stmt.Exec(0, video.Name, video.Url, video.LikeCount, video.DislikeCount)
	checkErr(err)
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
