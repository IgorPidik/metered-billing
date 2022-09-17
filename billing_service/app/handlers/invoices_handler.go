package handlers

import (
	"billing_service/app/models"

	"gorm.io/gorm"
)

type InvoicesHandler struct {
	DB *gorm.DB
}

func (i *InvoicesHandler) GetInvoices() ([]*models.Invoice, error) {
	var invoices []*models.Invoice
	if err := i.DB.Find(&invoices).Error; err != nil {
		return nil, err
	}

	return invoices, nil
}
