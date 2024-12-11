package repositories

import (
	"example/backend/internal/models"
	"time"
)

var transactions = []models.Transaction{
	{
		ID:          "12345",
		Type:        "despesa",
		Amount:      150.75,
		Description: "Compra de supermercado",
		Category:    "Alimentação",
		Date:        time.Now(),
	},
	{
		ID:          "12346",
		Type:        "despesa2",
		Amount:      10.75,
		Description: "Compra de supermercado",
		Category:    "Alimentação",
		Date:        time.Now(),
	},
}

func GetAllTransactions() []models.Transaction {
	return transactions
}
