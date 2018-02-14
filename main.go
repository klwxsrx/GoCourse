package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"gocourse/handlers"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	file, err := os.OpenFile("my.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
	}
	defer file.Close()

	rand.Seed(time.Now().UnixNano())

	killSignalChan := getKillSignalChan()
	srv := startServer(":8000")
	waitForKillSignal(killSignalChan)
	srv.Shutdown(context.Background())
}

func startServer(serverUrl string) *http.Server {
	log.WithFields(log.Fields{"url": serverUrl}).Info("starting the server")
	r := handlers.Router()
	srv := &http.Server{Addr: serverUrl, Handler: r}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
	return srv
}

func getKillSignalChan() chan os.Signal {
	osKillSignalChan := make(chan os.Signal, 1)
	signal.Notify(osKillSignalChan, os.Kill, os.Interrupt, syscall.SIGTERM)
	return osKillSignalChan
}

func waitForKillSignal(killSignalChan <-chan os.Signal) {
	killSignal := <-killSignalChan
	switch killSignal {
	case os.Interrupt:
		log.Print("got Interrupt")
	case syscall.SIGTERM:
		log.Print("got SIGTERM")
	}
}
