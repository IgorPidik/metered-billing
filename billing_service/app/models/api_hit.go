package models

import (
	"time"

	"github.com/google/uuid"
)

type APIHitKafka struct {
	UUID       uuid.UUID `json:"uuid"`
	CustomerID uint      `json:"customer_id"`
	ServiceID  uint      `json:"service_id"`
	Timestamp  time.Time `json:"timestamp"`
}

type APIHit struct {
	UUID        uuid.UUID `gorm:"primaryKey;type:uuid;unique"`
	CustomerID  uint
	ServiceID   uint
	Timestamp   time.Time
	InvoiceUUID *uuid.UUID
}
