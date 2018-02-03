package handlers

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func video(w http.ResponseWriter, _ *http.Request) {
	/*
		vars := mux.Vars(r)
		id := vars["ID"]
	*/

	item := VideoItem{
		VideoInfo{
			"d290f1ee-6c54-4b01-90e6-d701748f0851",
			"Black Retrospective Woman",
			15,
			"/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg",
		},
		"/content/d290f1ee-6c54-4b01-90e6-d701748f0851/index.mp4",
	}
	b, err := json.Marshal(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8;")
	_, err = io.WriteString(w, string(b))
	w.WriteHeader(http.StatusOK)
	if err != nil {
		log.WithField("err", err).Error("write response error")
	}
}
