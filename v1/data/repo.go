package data

import "github.com/shiva0705/goApi/v1/models"

var videos models.Videos

func init() {
	RepoCreateVideo(models.Video{Id: 1, Name: "Ultimate Dog Tease", Url: "https://www.youtube.com/watch?v=nGeKSiCQkPw", LikeCount: 0, DislikeCount: 0})
	RepoCreateVideo(models.Video{Id: 2, Name: "Dog Wants Kitty", Url: "https://www.youtube.com/watch?v=kI4yoXyb1_M", LikeCount: 0, DislikeCount: 0})
}

func GetVideos() models.Videos {
	return videos
}

func RepoCreateVideo(v models.Video) {
	videos = append(videos, v)
}

func RepoUpdateFeedback(f models.Feedback) {

	var updatedVideos models.Videos

	for _, v := range videos {
		if v.Id == f.VideoId {

			if f.Like == true {
				v.LikeCount++
			} else {
				v.DislikeCount++
			}
		}
		updatedVideos = append(updatedVideos, v)
	}
	videos = updatedVideos
}