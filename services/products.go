package services

import (
	Models "savannah/models"
)

//CRUD for products

func addProducts(item Models.Products) (Models.Products, error) {

	result := db.Create(item)

	if result.Error != nil {
		return product, result.Error
	}

	return product, nil

}

func getProductById(id uint) (Models.Products, error) {

	result := db.First(&product, id)

	if result.Error != nil {
		return product, result.Error
	}

	return product, nil

}

func getAllProducts() ([]Models.Products, error) {
	result := db.Find(&products)

	if result.Error != nil {
		return products, result.Error
	}

	return products, nil
}
