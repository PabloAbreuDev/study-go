package services

import (
	"example/backend/internal/models"
	"example/backend/internal/repositories"
)

func FetchTransactions() []models.Transaction {
	return repositories.GetAllTransactions()
}
