package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Invoice struct {
	UUID       uuid.UUID `gorm:"type:uuid;primaryKey;unique"`
	CustomerID uint
	Hits       []*APIHit
}

func (i *Invoice) BeforeCreate(tx *gorm.DB) (err error) {
	i.UUID = uuid.New()
	return nil
}
