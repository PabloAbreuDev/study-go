package services

import (
	"example/backend/internal/models"
	"example/backend/internal/repositories"
)

func FetchBudgets() []models.Budget {
	return repositories.GetAllBudgets()
}

func PostBudget(budget models.Budget) interface{} {
	return repositories.CreateBudget(budget)
}

func GetBudgetByID(id string) (*models.Budget, error) {
	return repositories.GetBudgetByID(id)
}

func UpdateBudget(id string, updatedData models.Budget) error {
	return repositories.UpdateBudget(id, updatedData)
}
