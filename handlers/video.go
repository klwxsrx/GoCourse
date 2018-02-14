package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"gocourse/view"
	"io"
	"net/http"
)

func video(w http.ResponseWriter, r *http.Request, vr VideoRepository) {
	vars := mux.Vars(r)
	videoKey := vars["id"]
	if videoKey == "" {
		http.Error(w, "missing id value", http.StatusBadRequest)
		return
	}

	video, err := vr.GetByKey(videoKey)
	if err != nil {
		log.WithField("err", err).Error("get video error")
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	if video == nil {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	item := view.GetVideoItemFromVideo(video)
	b, err := json.Marshal(item)
	if err != nil {
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
