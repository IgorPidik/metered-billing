package main

import (
	"billing_service/app/graph"
	"billing_service/app/graph/generated"
	"billing_service/app/handlers"
	"billing_service/app/invoicing"
	"billing_service/app/models"
	"billing_service/app/utils"
	"context"
	"encoding/json"
	"log"
	"net/http"

	// "time"

	// "github.com/go-co-op/gocron"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

const (
	topic          = "validated_hits"
	broker1Address = "kafka:9092"
)

func saveHit(db *gorm.DB, hit *models.APIHitKafka) error {
	return db.Create(&models.APIHit{UUID: hit.UUID, CustomerID: hit.CustomerID, ServiceID: hit.ServiceID, Timestamp: hit.Timestamp}).Error
}

func processHits(db *gorm.DB, reader *kafka.Reader) {
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		hit := &models.APIHitKafka{}
		if jsonErr := json.Unmarshal(msg.Value, hit); jsonErr != nil {
			log.Fatal(jsonErr)
		}

		if saveErr := saveHit(db, hit); saveErr != nil {
			log.Fatal(saveErr)
		}
	}
}

func main() {
	log.Println("Starting a billing service...")
	db, dbErr := utils.InitDB()

	if dbErr != nil {
		log.Fatal(dbErr)
	}
	// utils.CreateTestData(db)
	// Setup reader to consume validated hits
	// reader := kafka.NewReader(kafka.ReaderConfig{
	// 	Brokers:     []string{broker1Address},
	// 	Topic:       topic,
	// 	StartOffset: kafka.FirstOffset,
	// 	MinBytes:    5,
	// 	MaxBytes:    1e6,
	// 	MaxWait:     3 * time.Second,
	// })

	// defer reader.Close()

	// scheduler := gocron.NewScheduler(time.UTC)
	// scheduler.Every(1).Days().Do(doInvoicing)
	// scheduler.StartAsync()

	invoicing.DoInvoicing(db)

	// processHits(reader)
	hitsHandler := &handlers.HitsHandler{DB: db}
	invoicesHandler := &handlers.InvoicesHandler{DB: db}

	port := "8081"
	resolver := &graph.Resolver{
		HitsHandler:     hitsHandler,
		InvoicesHandler: invoicesHandler,
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
