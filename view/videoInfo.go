package view

import "gocourse/model"

type VideoInfo struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Duration  int    `json:"duration"`
	Thumbnail string `json:"thumbnail"`
	Status    int    `json:"status"`
}

func GetVideoInfoFromVideo(video *model.Video) *VideoInfo {
	return &VideoInfo{
		video.Key,
		video.Name,
		video.Duration,
		video.ThumbnailUrl,
		video.Status,
	}
}
