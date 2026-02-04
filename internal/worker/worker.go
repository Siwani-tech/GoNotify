package worker

import (
	"gonotify/internal/models"
	"gonotify/internal/queue"
	"log"
	"math/rand"
	"sync"
)

func StartWorker(id int, queue <-chan models.Notification, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Printf("Worker %d started\n", id)
	for notification := range queue {
		log.Printf("Worker %d processing: %s\n", id, notification.Message)

		if rand.Intn(2) == 0 {
			log.Println("Worker", id, "failed", notification.Message)
			handleFailure(notification)
			continue
		}
		log.Println("Worker", id, "sucess:", notification.Message)
	}
	log.Println("Worker", id, "stopped")
}

func handleFailure(notification models.Notification) {
	notification.Retries++
	if notification.Retries <= 3 {
		queue.NotificationQueue <- notification
		return
	}
	queue.DeadLetterQueue <- notification
}
