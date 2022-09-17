package handlers

import (
	"billing_service/app/models"

	"gorm.io/gorm"
)

type HitsHandler struct {
	DB *gorm.DB
}

func (h *HitsHandler) SaveHit(hit *models.APIHitKafka) (*models.APIHit, error) {
	storedHit := &models.APIHit{UUID: hit.UUID, CustomerID: hit.CustomerID, ServiceID: hit.ServiceID, Timestamp: hit.Timestamp}
	if err := h.DB.Create(storedHit).Error; err != nil {
		return nil, err
	}

	return storedHit, nil
}

func (h *HitsHandler) GetHits() ([]*models.APIHit, error) {
	var hits []*models.APIHit
	if err := h.DB.Find(&hits).Error; err != nil {
		return nil, err
	}

	return hits, nil
}
