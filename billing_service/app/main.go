package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

const (
	topic          = "validated_hits"
	broker1Address = "kafka:9092"
)

type APIHit struct {
	UUID       uuid.UUID `json:"uuid"`
	CustomerID uint      `json:"customer_id"`
	ServiceID  uint      `json:"service_id"`
	Timestamp  time.Time `json:"timestamp"`
}

func billHit(hit *APIHit) {

}

func runBilling(reader *kafka.Reader) {
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		hit := &APIHit{}
		if jsonErr := json.Unmarshal(msg.Value, hit); jsonErr != nil {
			log.Fatal(jsonErr)
		}

		billHit(hit)
	}
}

func main() {
	log.Println("Starting billing service...")

	// Setup reader to consume validated hits
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{broker1Address},
		Topic:       topic,
		StartOffset: kafka.FirstOffset,
		MinBytes:    5,
		MaxBytes:    1e6,
		MaxWait:     3 * time.Second,
	})
	defer reader.Close()

	runBilling(reader)
}
