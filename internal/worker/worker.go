package worker

import (
	"fmt"
	"gonotify/internal/models"
	"gonotify/internal/queue"
	"math/rand"
	"sync"
)

func StartWorker(id int, queue <-chan models.Notification, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started\n", id)
	for notification := range queue {
		fmt.Printf("Worker %d processing: %s\n", id, notification.Message)

		if rand.Intn(2) == 0 {
			fmt.Println("Worker", id, "failed", notification.Message)
			handleFailure(notification)
			continue
		}
		fmt.Println("Worker", id, "sucess:", notification.Message)
	}
	fmt.Println("Worker", id, "stopped")
}

func handleFailure(notification models.Notification) {
	notification.Retries++
	if notification.Retries <= 3 {
		queue.NotificationQueue <- notification
		return
	}
	queue.DeadLetterQueue <- notification
}
