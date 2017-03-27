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
	mock.ExpectExec("Insert Into person").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	var video = models.Video{Id: 1, Name: "Ultimate Dog Tease", Url: "https://www.youtube.com/watch?v=nGeKSiCQkPw", LikeCount: 0, DislikeCount: 0}
	if err = addVideo(video); err != nil {
		t.Errorf("error was not expected while adding video: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfiled expectations: %s", err)
	}
}
