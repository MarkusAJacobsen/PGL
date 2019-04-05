package main

import (
	"github.com/MarkusAJacobsen/pgl/internal"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		logrus.Panic("Port not sat")
	}

	r := mux.NewRouter()
	rr := r.PathPrefix("/report").Subrouter()
	rr.HandleFunc("/error", internal.HandleError).Methods(http.MethodPost)
	rr.HandleFunc("/traffic", internal.HandleTraffic).Methods(http.MethodPost)

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	printStartUpMsg(port)
	logrus.Fatal(srv.ListenAndServe())
}

func printStartUpMsg(port string) {
	logrus.Infof("Starting up PGL on port: %s", port)
}
