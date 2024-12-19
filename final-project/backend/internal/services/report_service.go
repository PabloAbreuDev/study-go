package services

import (
	"example/backend/internal/models"
	"example/backend/internal/repositories"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateReport(startDate, endDate time.Time) ([]models.Report, error) {
	transactions, err := repositories.GetTransactionsByDateRange(startDate, endDate)
	if err != nil {
		return nil, err
	}

	categories := repositories.GetAllCategories()
	if err != nil {
		return nil, err
	}

	categoryMap := make(map[primitive.ObjectID]string)

	for _, category := range categories {
		objectID, err := primitive.ObjectIDFromHex(category.ID)
		if err != nil {
			continue
		}
		categoryMap[objectID] = category.Name
	}

	reportMap := make(map[string]*models.Report)
	for _, transaction := range transactions {
		categoryName, exists := categoryMap[transaction.CategoryID]

		if !exists {
			categoryName = "Unknown" // Categoria não encontrada
		}

		if _, exists := reportMap[transaction.CategoryID.String()]; !exists {
			reportMap[transaction.CategoryID.String()] = &models.Report{
				CategoryName: categoryName,
				TotalAmount:  0,
				Transactions: []models.Transaction{},
			}
		}

		reportMap[transaction.CategoryID.String()].TotalAmount += transaction.Amount
		reportMap[transaction.CategoryID.String()].Transactions = append(reportMap[transaction.CategoryID.String()].Transactions, transaction)
	}

	var reports []models.Report
	for _, report := range reportMap {
		reports = append(reports, *report)
	}

	return reports, nil
}
