package model

import (
	"math/rand"
)

type Video struct {
	id           int
	Key          string
	Name         string
	Status       int
	Duration     int
	ThumbnailUrl string
	Url          string
}

const (
	videoKeyAvailableLetters = "abcdefghijklmnopqrstuvwxyz0123456789"

	StatusError  = 1
	StatusUpload = 2
	StatusReady  = 3
)

func NewVideo(name string, status int) *Video {
	return &Video{0, generateVideoKey(), name, status, 0, "", ""}
}

func generateVideoKey() string {
	var availableLetters = []rune(videoKeyAvailableLetters)
	var key = make([]rune, 36)
	for i := range key {
		if i == 8 || i == 13 || i == 18 || i == 23 {
			key[i] = '-'
		} else {
			key[i] = availableLetters[rand.Intn(len(availableLetters))]
		}
	}
	return string(key)
}
