package server

import (
	"net/http"

	"gonotify/internal/handlers"
)

func StartServer() *http.Server {

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GoNotify is running"))
	})

	mux.HandleFunc("/notify", handlers.NotifyHandler)

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}
