package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
)

type Tweet struct {
	Country string `json:"country"`
	Weather string `json:"weather"`
}

func main() {
	brokers := strings.Split(os.Getenv("KAFKA_BROKERS"), ",") // kafka:9092
	topic := os.Getenv("KAFKA_TOPIC")                         // message
	groupID := os.Getenv("KAFKA_GROUP")                       // g1
	redisAddr := os.Getenv("REDIS_ADDR")                      // redis-master.default.svc.cluster.local:6379

	// Conecta a Redis
	rdb := redis.NewClient(&redis.Options{Addr: redisAddr})
	defer rdb.Close()

	// Crea el reader de Kafka que siempre arranca desde el principio
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     brokers,
		Topic:       topic,
		GroupID:     groupID,
		StartOffset: kafka.FirstOffset,
		MinBytes:    10e3,
		MaxBytes:    10e6,
	})
	defer reader.Close()

	log.Println("Kafka consumer started...")

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v\n", err)
			time.Sleep(time.Second)
			continue
		}
		var t Tweet
		if err := json.Unmarshal(m.Value, &t); err != nil {
			log.Printf("Error unmarshalling JSON: %v\n", err)
			continue
		}
		if err := rdb.HIncrBy(context.Background(), "tweets", t.Country, 1).Err(); err != nil {
			log.Printf("Error writing to Redis: %v\n", err)
			continue
		}
		log.Printf("Processed message for country %s\n", t.Country)
	}
}
