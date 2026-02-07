# GoNotify 

GoNotify is a lightweight notification processing service built with Go.
It accepts notification requests over HTTP and processes them asynchronously using goroutines, channels, and a worker pool.
---

##  Architecture Overview

```
Client
  |
  |  POST /notify
  v
HTTP Handler
  |
  |  pushes notification
  v
Buffered Channel (Queue)
  |
  |  fan-out
  v
Worker Pool (goroutines)
  |        |
  | success| failure
  v        v
Done     Dead-Letter Queue
            |
            v
        DLQ Worker (logs / retry-ready)
```

---

##  Core Concepts Used

* **Goroutines** – lightweight concurrent execution units
* **Channels** – safe communication between goroutines
* **Worker Pool** – multiple workers consuming from one queue
* **Dead-Letter Queue (DLQ)** – failed notifications are preserved, not lost
* **Environment-based configuration** – runtime tuning without code changes
* **Graceful shutdown** – clean exit on SIGTERM / Ctrl+C

---

##  Project Structure

```
cmd/
  main.go                # App entry point & orchestration

internal/
  config/                # Env config helpers
  handlers/              # HTTP request handlers
  models/                # Domain models
  queue/                 # Notification & DLQ channels
  worker/                # Worker & DLQ worker logic
  server/                # HTTP server setup
```
---

##  Running the Project

### 1 Clone the repo

```bash
git clone [https://github.com/Siwani-tech/gonotify.git](https://github.com/Siwani-tech/GoNotify.git)
cd gonotify
```

### 2 Run the server

```bash
go run main.go
```

Server starts on:

```
http://localhost:8080
```

---

##  API Usage

### Health Check

```
GET /health
```

Response:

```
GoNotify is running
```

---

### Send Notifications

```
POST /notify
Content-Type: application/json
```

#### Example Request

```json
[
  { "id": "1", "type": "email", "message": "Hello" },
  { "id": "2", "type": "sms", "message": "Hi" },
  { "id": "3", "type": "email", "message": "Again" }
]
```

#### Example Response

```
202 Accepted
All notifications queued
```






