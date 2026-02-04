package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"gonotify/internal/config"
	"gonotify/internal/queue"
	"gonotify/internal/server"
	"gonotify/internal/worker"
)

func main() {
	numWorkers := config.GetEnvInt("NUM_WORKERS", 3)
	queueSize := config.GetEnvInt("QUEUE_SIZE", 100)
	queue.Init(queueSize)

	go worker.StartDLQWorker(queue.DeadLetterQueue)
	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker.StartWorker(i, queue.NotificationQueue, &wg)
	}
	srv := server.StartServer()

	go func() {
		log.Println("Server started on :8080")
		if err := srv.ListenAndServe(); err != nil && err.Error() != "http: Server closed" {
			log.Fatal(err)
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	close(queue.NotificationQueue)
	wg.Wait()

	log.Println("Shutdown complete")
}
