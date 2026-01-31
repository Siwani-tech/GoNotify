package worker

import (
	"fmt"
	"gonotify/internal/models"
	"time"
)

func StartWorker(id int, queue <-chan models.Notification) {
	fmt.Printf("Worker %d started\n", id)
	for notification := range queue {
		fmt.Printf("Worker %d processing: %s\n", id, notification.Message)
		time.Sleep(2 * time.Second)
		fmt.Printf("Worker %d finished\n", id)
	}
}
