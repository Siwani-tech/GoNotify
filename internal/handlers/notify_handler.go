package handlers

import (
	"encoding/json"
	"net/http"

	"gonotify/internal/models"
	"gonotify/internal/queue"
)

func NotifyHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var notifications []models.Notification
	err := json.NewDecoder(r.Body).Decode(&notifications)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	failed := 0
	for _, notification := range notifications {
		select {
		case queue.NotificationQueue <- notification:
		default:
			failed++
		}
	}

	if failed > 0 {
		w.WriteHeader(http.StatusPartialContent)
		w.Write([]byte("Some notifications could not be queued"))
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("All notifications queued"))
}
