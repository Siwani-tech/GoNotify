package worker

import (
	"gonotify/internal/models"
	"log"
)

func StartDLQWorker(dlq <-chan models.Notification) {
	for notification := range dlq {
		log.Println("DLQ notification:", notification.Message)
	}
}
