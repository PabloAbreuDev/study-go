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

	categoryCollection := repositories.GetAllCategories()

	categoryMap := make(map[primitive.ObjectID]models.Category)
	for _, category := range categoryCollection {
		objectID, err := primitive.ObjectIDFromHex(category.ID)
		if err != nil {
			return nil, err
		}
		categoryMap[objectID] = category
	}

	// Group transactions by category
	reportMap := make(map[primitive.ObjectID]*models.Report)
	for _, transaction := range transactions {
		categoryName := categoryMap[transaction.CategoryID]

		// Initialize if the category isn't present in the report map
		if _, exists := reportMap[transaction.CategoryID]; !exists {
			reportMap[transaction.CategoryID] = &models.Report{
				CategoryName: categoryName.Name,
				TotalAmount:  0,
				Transactions: []models.Transaction{},
			}
		}

		// Update the report for the category
		reportMap[transaction.CategoryID].TotalAmount += transaction.Amount
		reportMap[transaction.CategoryID].Transactions = append(reportMap[transaction.CategoryID].Transactions, transaction)
	}

	// Convert the map to a slice
	var reports []models.Report
	for _, report := range reportMap {
		reports = append(reports, *report)
	}

	return reports, nil
}
