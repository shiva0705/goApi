package models

type Feedback struct {
	VideoId int  `json:"videoId"`
	Like    bool `json:"like"`
}
