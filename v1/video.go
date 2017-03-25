package main

type Video struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Url          string `json:"url"`
	LikeCount    int    `json:"likeCount"`
	DislikeCount int    `json:"dislikeCount"`
}

type Videos []Video
