package services

import (
	"example/backend/internal/models"
	"example/backend/internal/repositories"
)

func FetchCategories() []models.Category {
	return repositories.GetAllCategories()
}

func PostCategory(category models.Category) interface{} {
	return repositories.CreateCategory(category)
}

func GetCategoryByID(id string) (*models.Category, error) {
	return repositories.GetCategoryByID(id)
}

func UpdateCategory(id string, updatedData models.Category) error {
	return repositories.UpdateCategory(id, updatedData)
}
