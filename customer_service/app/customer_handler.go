package main

import (
	"customer_service/app/models"
	"errors"

	"gorm.io/gorm"
)

var InvalidCustomerIdErr = errors.New("invalid customer ID")

type CustomerHandler struct {
	DB *gorm.DB
}

func (c *CustomerHandler) CreateCustomer(name string) (*models.Customer, error) {
	customer := &models.Customer{Name: name}
	if err := c.DB.Create(customer).Error; err != nil {
		return nil, err
	}

	return customer, nil
}

func (c *CustomerHandler) CreateService(customerId uint, name string) (*models.Service, error) {
	if err := c.DB.First(&models.Customer{Model: gorm.Model{ID: customerId}}).Error; err != nil {
		return nil, InvalidCustomerIdErr
	}

	service := &models.Service{Name: name, CustomerID: customerId}
	if err := c.DB.Create(service).Error; err != nil {
		return nil, err
	}

	return service, nil
}

func (c *CustomerHandler) ServiceBelongsToCustomer(serviceId uint, customerId uint) (bool, error) {
	service := &models.Service{
		Model:      gorm.Model{ID: serviceId},
		CustomerID: customerId,
	}

	if err := c.DB.First(service).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
