package services

import (
	Models "savannah/models"
)

func addCategory(item Models.Categories) (Models.Categories, error) {

	result := db.Create(item)

	if result.Error != nil {
		return category, result.Error
	}

	return category, nil

}

func getCategoryById(id uint) (Models.Categories, error) {

	result := db.First(&category, id)

	if result.Error != nil {
		return category, result.Error
	}

	return category, nil

}

func getAllCategories() ([]Models.Categories, error) {

	result := db.Find(&categories)

	if result.Error != nil {
		return categories, result.Error
	}

	return categories, nil
}
