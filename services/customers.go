package services

import (
	Models "savannah/models"
)

func addCustomer(item Models.Customers) (Models.Customers, error) {

	result := db.Create(item)

	if result.Error != nil {
		return customer, result.Error
	}

	return customer, nil

}

func getCustomerById(id uint) (Models.Customers, error) {

	result := db.First(&customer, id)

	if result.Error != nil {
		return customer, result.Error
	}

	return customer, nil

}

func getAllCustomers() ([]Models.Customers, error) {

	result := db.Find(&customers)

	if result.Error != nil {
		return customers, result.Error
	}

	return customers, nil
}
