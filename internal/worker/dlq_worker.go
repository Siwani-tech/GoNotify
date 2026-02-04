package worker

import (
	"fmt"

	"gonotify/internal/models"
)

func StartDLQWorker(dlq <-chan models.Notification) {
	for notification := range dlq {
		fmt.Println("DLQ notification:", notification.Message)
	}
}
