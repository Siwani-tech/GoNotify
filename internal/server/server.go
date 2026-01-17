package server

import (
	"fmt"
	"net/http"

	"gonotify/internal/handlers"
)

func StartServer() {
	fmt.Println("Server started on port 8080")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GoNotify is running"))
	})

	http.HandleFunc("/notify", handlers.NotifyHandler)
	http.ListenAndServe(":8080", nil)
}
