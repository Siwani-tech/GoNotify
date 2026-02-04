package queue

import "gonotify/internal/models"

var NotificationQueue = make(chan models.Notification, 100)
var DeadLetterQueue = make(chan models.Notification, 100)
