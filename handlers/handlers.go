package handlers

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"gocourse/model"
	"net/http"
)

type VideoRepository interface {
	GetByKey(string) (*model.Video, error)
	GetAll() ([]*model.Video, error)
	Save(*model.Video) error
}

func Router() http.Handler {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()

	s.HandleFunc("/list", videosHandler(list)).Methods(http.MethodGet)
	s.HandleFunc("/video/{id:[\\w-]+}", videosHandler(video)).Methods(http.MethodGet)
	s.HandleFunc("/video", videosHandler(uploadVideo)).Methods(http.MethodPost)

	return logMiddleware(r)
}

func videosHandler(f func(http.ResponseWriter, *http.Request, VideoRepository)) func(http.ResponseWriter, *http.Request) {
	vr := model.NewVideoRepository()
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r, vr)
	}
}

func logMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"method":     r.Method,
			"url":        r.URL,
			"remoteAddr": r.RemoteAddr,
			"userAgent":  r.UserAgent(),
		}).Info("got a new request")
		h.ServeHTTP(w, r)
	})
}
