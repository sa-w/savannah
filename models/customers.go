package models

import "gorm.io/gorm"

type Customers struct {
	gorm.Model
	Name  string
	Email *string
}
