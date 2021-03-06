package data

import (
	"testing"

	"github.com/shiva0705/goApi/v1/models"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestAddVideo(t *testing.T) {
	db, mock, err := sqlmock.New()
	checkErr(err)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("Insert").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	var video = models.Video{Id: 1, Name: "Ultimate Dog Tease", Url: "https://www.youtube.com/watch?v=nGeKSiCQkPw", LikeCount: 0, DislikeCount: 0}
	if err = addVideo(db, video); err != nil {
		t.Errorf("error was not expected while adding video: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfiled expectations: %s", err)
	}
}

func TestGetVideos(t *testing.T) {
	db, mock, err := sqlmock.New()
	checkErr(err)
	defer db.Close()

	var columns = []string{"id", "name", "url", "like_count", "dislike_count"}
	mock.ExpectQuery("Select").WillReturnRows(sqlmock.NewRows(columns).AddRow(0, "t", "t", 0, 0))

	var videos models.Videos
	videos = GetVideos(db)

	if videos == nil {
		t.Errorf("error as no videos were returned")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfiled expectations: %s", err)
	}

}

func TestUpdateFeedback(t *testing.T) {
	db, mock, err := sqlmock.New()
	checkErr(err)
	defer db.Close()

	var columns = []string{"id", "name", "url", "like_count", "dislike_count"}
	mock.ExpectQuery("Select").WillReturnRows(sqlmock.NewRows(columns).AddRow(0, "t", "t", 0, 0))
	mock.ExpectPrepare("Update")
	mock.ExpectExec("Update").WithArgs(1, 0, 0).WillReturnResult(sqlmock.NewResult(0, 1))

	var feedback = models.Feedback{VideoId: 0, Like: true}
	if err := UpdateFeedback(db, feedback); err != nil {
		t.Errorf("error while updating feedback: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfiled expectations: %s", err)
	}

}
