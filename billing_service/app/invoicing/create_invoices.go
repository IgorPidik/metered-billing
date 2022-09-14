package invoicing

import (
	"billing_service/app/models"
	"log"

	"gorm.io/gorm"
)

type customer struct {
	CustomerID uint
}

func invoiceCustomer(db *gorm.DB, customerId uint) error {
	log.Printf("billing %v\n", customerId)
	invoice := &models.Invoice{CustomerID: customerId}
	if err := db.Create(invoice).Error; err != nil {
		return err
	}

	if invoiceErr := db.Model(&models.APIHit{}).Where("customer_id = ?", customerId).Update("invoice_uuid", invoice.UUID).Error; invoiceErr != nil {
		db.Delete(invoice)
		return invoiceErr
	}

	return nil
}

func findCustomers(db *gorm.DB) ([]customer, error) {
	var customers []customer
	err := db.Model(&models.APIHit{}).Distinct("customer_id").Where("invoice_uuid IS NULL").Find(&customers).Error
	return customers, err
}

func DoInvoicing(db *gorm.DB) error {
	customers, customersErr := findCustomers(db)
	if customersErr != nil {
		return customersErr
	}

	for _, customer := range customers {
		err := invoiceCustomer(db, customer.CustomerID)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
