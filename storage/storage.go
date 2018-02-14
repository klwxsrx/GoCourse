package storage

import (
	"errors"
	"gocourse/model"
	"io"
	"mime/multipart"
	"os"
)

const videoUploadPath = "content/"
const videoFileName = "index.mp4"

func UploadVideo(video *model.Video, file multipart.File) (string, string, error) {
	folderPath := getFolderPath(video.Key)
	err := createFolder(folderPath)
	if err != nil {
		return "", "", errors.New("failed to create path: " + err.Error())
	}

	contentPath := getContentPath(video.Key)
	err = saveVideo(file, contentPath)
	if err != nil {
		return "", "", errors.New("failed to save file: " + err.Error())
	}

	return "/" + contentPath, getThumbnailUrl(video.Key), nil
}

func createFolder(folderPath string) error {
	return os.Mkdir(folderPath, os.ModeDir)
}

func saveVideo(file multipart.File, contentPath string) error {
	fileTo, _ := os.OpenFile(contentPath, os.O_RDWR|os.O_CREATE, 0666)
	_, err := io.Copy(fileTo, file)
	return err
}

func getFolderPath(key string) string {
	return videoUploadPath + key
}

func getContentPath(key string) string {
	return getFolderPath(key) + "/" + videoFileName
}

func getThumbnailUrl(_ string) string {
	return ""
}
