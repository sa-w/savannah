package services

import (
	Models "savannah/models"
)

func addOrder(item Models.Orders) (Models.Orders, error) {

	result := db.Create(item)

	if result.Error != nil {
		return order, result.Error
	}

	return order, nil

}

func getOrderById(id uint) (Models.Orders, error) {

	result := db.First(&order, id)

	if result.Error != nil {
		return order, result.Error
	}

	return order, nil

}

func getAllOrders() ([]Models.Orders, error) {

	result := db.Find(&orders)

	if result.Error != nil {
		return orders, result.Error
	}

	return orders, nil
}
