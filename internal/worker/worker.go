package worker

import (
	"fmt"
	"gonotify/internal/models"
	"sync"
)

func StartWorker(id int, queue <-chan models.Notification, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started\n", id)
	for notification := range queue {
		fmt.Printf("Worker %d processing: %s\n", id, notification.Message)
	}
	fmt.Println("Worker", id, "stopped")
}
