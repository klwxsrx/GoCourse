package view

import "gocourse/model"

type VideoItem struct {
	VideoInfo
	Url string `json:"url"`
}

func GetVideoItemFromVideo(video *model.Video) *VideoItem {
	return &VideoItem{*GetVideoInfoFromVideo(video), video.Url}
}
