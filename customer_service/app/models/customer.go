package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name     string
	Services []Service
}

type Service struct {
	gorm.Model
	Name       string
	CustomerID uint
}
