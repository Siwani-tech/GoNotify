package handlers

import (
	"encoding/json"
	"net/http"

	"gonotify/internal/models"
)

func NotifyHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var notification models.Notification

	err := json.NewDecoder(r.Body).Decode(&notification)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request body"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Notification received"))

}
