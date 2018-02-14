package handlers

import (
	log "github.com/sirupsen/logrus"
	"gocourse/model"
	"gocourse/storage"
	"net/http"
)

func uploadVideo(w http.ResponseWriter, r *http.Request, vr VideoRepository) {
	file, fileHeader, err := r.FormFile("file[]")
	if err != nil || fileHeader.Header.Get("Content-Type") != "video/mp4" {
		http.Error(w, "invalid request parameters", http.StatusBadRequest)
		return
	}

	video := model.NewVideo(fileHeader.Filename, model.StatusUpload)

	if err := vr.Save(video); err != nil {
		log.WithField("err", err).Error("save video to database failed")
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	url, thumbnailUrl, err := storage.UploadVideo(video, file)
	if err != nil {
		log.WithField("err", err).Error("upload video failed")
		http.Error(w, "", http.StatusInternalServerError)

		video.Status = model.StatusError
		if vr.Save(video); err != nil {
			log.WithField("err", err).Error("update video status to error in database failed")
		}

		return
	}

	video.Url = url
	video.ThumbnailUrl = thumbnailUrl
	video.Status = model.StatusReady

	if err = vr.Save(video); err != nil {
		log.WithField("err", err).Error("update video status to ready in database failed")
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
