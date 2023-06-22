package server

import (
	"net/http"
	"os"
	"time"
)

const (
	readHeaderTimeout = 10 * time.Second
)

func NewServer(
	rootHandler http.Handler,
) *http.Server {
	return &http.Server{
		Addr:              ":" + os.Getenv("PORT"),
		Handler:           rootHandler,
		ReadHeaderTimeout: readHeaderTimeout,
	}
}
