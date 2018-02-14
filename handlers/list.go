package handlers

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"gocourse/view"
	"io"
	"net/http"
)

func list(w http.ResponseWriter, _ *http.Request, r VideoRepository) {
	videos, err := r.GetAll()
	if err != nil {
		log.WithField("err", err).Error("get videos error")
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	var items []*view.VideoInfo
	for _, video := range videos {
		items = append(items, view.GetVideoInfoFromVideo(video))
	}

	b, err := json.Marshal(items)
	if err != nil {
		log.WithField("err", err).Error("json marshall error")
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8;")
	_, err = io.WriteString(w, string(b))
	w.WriteHeader(http.StatusOK)
	if err != nil {
		log.WithField("err", err).Error("write response error")
	}
}
