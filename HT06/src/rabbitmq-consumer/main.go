package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
	"github.com/streadway/amqp"
)

type Tweet struct {
	Country string `json:"country"`
	Weather string `json:"weather"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}

func main() {
	// Leer configuración desde variables de entorno
	amqpURL := os.Getenv("RABBITMQ_URL")
	queue := os.Getenv("RABBITMQ_QUEUE")
	valkeyAddr := os.Getenv("VALKEY_ADDR")
	valkeyPwd := os.Getenv("VALKEY_PASSWORD")

	// Conectar a RabbitMQ
	conn, err := amqp.Dial(amqpURL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declarar la cola si no existe
	_, err = ch.QueueDeclare(
		queue, // name
		true,  // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,   // args
	)
	failOnError(err, "Failed to declare queue")

	// Registrar el consumidor
	msgs, err := ch.Consume(
		queue, // queue
		"",    // consumer
		true,  // autoAck
		false, // exclusive
		false, // noLocal
		false, // noWait
		nil,   // args
	)
	failOnError(err, "Failed to register a consumer")

	// Conectar a Valkey (protocolo Redis)
	vk := redis.NewClient(&redis.Options{
		Addr:     valkeyAddr,
		Password: valkeyPwd,
	})
	defer vk.Close()

	log.Println("RabbitMQ consumer started...")
	for d := range msgs {
		var t Tweet
		if err := json.Unmarshal(d.Body, &t); err != nil {
			log.Printf("Invalid payload: %v\n", err)
			continue
		}
		// Incrementar contador por país en Valkey
		if err := vk.HIncrBy(context.Background(), "tweets_vk", t.Country, 1).Err(); err != nil {
			log.Printf("Error saving to Valkey: %v\n", err)
		}
	}
}
