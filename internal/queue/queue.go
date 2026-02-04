package queue

import "gonotify/internal/models"

var NotificationQueue chan models.Notification
var DeadLetterQueue chan models.Notification

func Init(queueSize int) {
	NotificationQueue = make(chan models.Notification, queueSize)
	DeadLetterQueue = make(chan models.Notification, queueSize)
}
