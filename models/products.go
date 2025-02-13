package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	Name      string
	Quantity  int
	UnitPrice float32
}
