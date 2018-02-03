package handlers

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func list(w http.ResponseWriter, _ *http.Request) {
	items := []VideoInfo{
		{
			"d290f1ee-6c54-4b01-90e6-d701748f0851",
			"Black Retrospective Woman",
			15,
			"/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg",
		},
		{
			"sldjfl34-dfgj-523k-jk34-5jk3j45klj34",
			"Go Rally TEASER-HD",
			41,
			"/content/sldjfl34-dfgj-523k-jk34-5jk3j45klj34/screen.jpg",
		},
		{
			"hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345",
			"Танцор",
			92,
			"/content/hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345/screen.jpg",
		},
	}

	b, err := json.Marshal(items)
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
