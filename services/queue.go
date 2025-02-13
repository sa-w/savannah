package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

// Add message to queue
func AddMessage(message string) {
	msg := message
	if msg != "" {
		client.RPush(ctx, "order_queue", msg)
		fmt.Println("Queued:", msg)
		time.Sleep(500 * time.Millisecond)
	}
}

// process messages in queue
func ProcessMessage() {
	for {
		msg, err := client.LPop(ctx, "order_queue").Result()
		if err == redis.Nil {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Processing:", msg)
		time.Sleep(500 * time.Millisecond) // Simulate processing delay
	}
}
