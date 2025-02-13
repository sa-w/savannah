package models

import "gorm.io/gorm"

type Orders struct {
	gorm.Model
	Quantity int
}
