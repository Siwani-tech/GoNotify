package server

import (
	"fmt"
	"net/http"

	"gonotify/internal/handlers"
	"gonotify/internal/queue"
	"gonotify/internal/worker"
)

func StartServer() {
	fmt.Println("Server started on port 8080")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GoNotify is running"))
	})

	http.HandleFunc("/notify", handlers.NotifyHandler)
	for i := 1; i <= 3; i++ {
		go worker.StartWorker(i, queue.NotificationQueue)
	}
	http.ListenAndServe(":8080", nil)
}
