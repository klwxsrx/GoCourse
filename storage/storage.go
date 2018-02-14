package storage

import (
	"errors"
	"gocourse/model"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

const videoUploadPath = "content"
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
	return os.Mkdir(folderPath, os.ModePerm)
}

func saveVideo(file multipart.File, contentPath string) error {
	fileTo, err := os.OpenFile(contentPath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	_, err = io.Copy(fileTo, file)
	return err
}

func getFolderPath(key string) string {
	return filepath.Join(videoUploadPath, key)
}

func getContentPath(key string) string {
	return filepath.Join(getFolderPath(key), videoFileName)
}

func getThumbnailUrl(_ string) string {
	return ""
}
