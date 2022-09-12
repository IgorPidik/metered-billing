package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

const (
	topic          = "consumed_hits"
	broker1Address = "kafka:9092"
)

type APIHit struct {
	CustomerID uint `json:"customer_id"`
	ServiceID  uint `json:"service_id"`
}

type APIHitKafkaMessage struct {
	APIHit
	UUID      uuid.UUID `json:"uuid"`
	Timestamp time.Time `json:"timestamp"`
}

func writeHitToKafka(hit *APIHit, writer *kafka.Writer) error {
	message := &APIHitKafkaMessage{
		APIHit:    *hit,
		UUID:      uuid.New(),
		Timestamp: time.Now(),
	}

	messageBytes, jsonErr := json.Marshal(message)
	if jsonErr != nil {
		return jsonErr
	}

	log.Println(string(messageBytes))

	return writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(fmt.Sprintf("%v-%v", hit.CustomerID, hit.ServiceID)),
		Value: messageBytes,
	})
}

func apiUsed(w http.ResponseWriter, r *http.Request, writer *kafka.Writer) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	hit := &APIHit{}
	err := json.NewDecoder(r.Body).Decode(hit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if writingErr := writeHitToKafka(hit, writer); writingErr != nil {
		log.Fatal(writingErr)
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	log.Println("Starting consumer...")
	writer := &kafka.Writer{
		Addr:  kafka.TCP(broker1Address),
		Topic: topic,
		// BatchSize: 1,
	}

	defer writer.Close()

	http.HandleFunc("/api-used", func(w http.ResponseWriter, r *http.Request) {
		apiUsed(w, r, writer)
	})

	log.Fatal(http.ListenAndServe(":3333", nil))
}
