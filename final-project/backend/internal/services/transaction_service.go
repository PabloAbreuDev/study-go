package services

import (
	"example/backend/internal/models"
	"example/backend/internal/repositories"
)

func FetchTransactions() []models.Transaction {
	return repositories.GetAllTransactions()
}

func PostTransactions(transaction models.Transaction) interface{} {
	return repositories.CreateTransaction(transaction)
}

func GetTransactionByID(id string) (*models.Transaction, error) {
	return repositories.GetTransactionByID(id)
}

func UpdateTransaction(id string, updatedData models.Transaction) error {
	return repositories.UpdateTransaction(id, updatedData)
}

func DeleteTransaction(id string) error {
	return repositories.DeleteTransaction(id)
}
