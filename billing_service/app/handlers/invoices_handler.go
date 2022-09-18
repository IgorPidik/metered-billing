package handlers

import (
	"billing_service/app/models"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InvoicesHandler struct {
	DB *gorm.DB
}

func (i *InvoicesHandler) GetInvoices() ([]*models.Invoice, error) {
	var invoices []*models.Invoice
	if err := i.DB.Preload("Hits").Find(&invoices).Error; err != nil {
		return nil, err
	}

	return invoices, nil
}

func (i *InvoicesHandler) GetInvoice(uuid uuid.UUID) (*models.Invoice, error) {
	var invoice *models.Invoice
	if dbErr := i.DB.Where("uuid = ?", uuid).Preload("Hits").First(&invoice).Error; dbErr != nil {
		if errors.Is(dbErr, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, dbErr
	}

	return invoice, nil
}

func (i *InvoicesHandler) GetInvoicesForCustomerId(id uint) ([]*models.Invoice, error) {
	var invoices []*models.Invoice
	if dbErr := i.DB.Where("customer_id = ?", id).Preload("Hits").Find(&invoices).Error; dbErr != nil {
		return nil, dbErr
	}

	return invoices, nil
}
