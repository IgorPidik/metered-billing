package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	pb "customer_service_proto"

	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	readerTopic    = "consumed_hits"
	writerTopic    = "validated_hits"
	broker1Address = "kafka:9092"
)

type APIHitKafkaMessage struct {
	CustomerID uint `json:"customer_id"`
	ServiceID  uint `json:"service_id"`
}

func validateHit(hit *APIHitKafkaMessage, cs pb.CustomerServiceClient) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	customerDetailsResponse, err := cs.CheckCustomerDetailsValidity(ctx, &pb.CustomerDetails{CustomerID: uint32(hit.CustomerID), ServiceID: uint32(hit.ServiceID)})
	if err != nil {
		return false, err
	}

	return customerDetailsResponse.Valid, nil
}

func runValidation(reader *kafka.Reader, writer *kafka.Writer, cs pb.CustomerServiceClient) {
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		hit := &APIHitKafkaMessage{}
		if jsonErr := json.Unmarshal(msg.Value, hit); jsonErr != nil {
			log.Fatal(jsonErr)
		}

		validHit, validationErr := validateHit(hit, cs)
		if validationErr != nil {
			log.Fatal(validationErr)
		}

		if validHit {
			if writingErr := writer.WriteMessages(context.Background(), kafka.Message{Key: msg.Key, Value: msg.Value}); writingErr != nil {
				log.Fatal(writingErr)
			}
		}
	}
}

func main() {
	log.Println("Starting validator...")
	// Set up a connection to the server.
	conn, err := grpc.Dial("customer_service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	cs := pb.NewCustomerServiceClient(conn)

	// Setup reader to consume hits
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{broker1Address},
		Topic:       readerTopic,
		StartOffset: kafka.FirstOffset,
		GroupID:     "validator_group",
		MinBytes:    5,
		MaxBytes:    1e6,
		MaxWait:     3 * time.Second,
	})

	defer reader.Close()

	// Setup writer to forward validated hits
	writer := &kafka.Writer{
		Addr:  kafka.TCP(broker1Address),
		Topic: writerTopic,
	}
	defer writer.Close()

	runValidation(reader, writer, cs)
}
