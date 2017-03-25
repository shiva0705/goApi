package main

var videos Videos

func init() {
	RepoCreateVideo(Video{Id: 1, Name: "Ultimate Dog Tease", Url: "https://www.youtube.com/watch?v=nGeKSiCQkPw", LikeCount: 0, DislikeCount: 0})
	RepoCreateVideo(Video{Id: 2, Name: "Dog Wants Kitty", Url: "https://www.youtube.com/watch?v=kI4yoXyb1_M", LikeCount: 0, DislikeCount: 0})
}

func RepoCreateVideo(v Video) {
	videos = append(videos, v)
}

func RepoUpdateFeedback(f Feedback) {

	for i, v := range videos {
		if v.Id == f.VideoId {

			if f.Like == true {
				v.LikeCount++
			} else {
				v.DislikeCount++
			}

			videos[i] = v
		}
	}
}
