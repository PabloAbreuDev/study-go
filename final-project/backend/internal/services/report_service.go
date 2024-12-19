package services

import (
	"example/backend/internal/models"
	"example/backend/internal/repositories"
	"fmt"
	"strings"
	"time"

	"github.com/jung-kurt/gofpdf"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateReportByType(startDate, endDate time.Time) (map[string][]models.Report, float64, error) {
	transactions, err := repositories.GetTransactionsByDateRange(startDate, endDate)
	if err != nil {
		return nil, 0, err
	}

	categories := repositories.GetAllCategories()
	if err != nil {
		return nil, 0, err
	}

	// Mapear categorias
	categoryMap := make(map[primitive.ObjectID]string)
	for _, category := range categories {
		categoryID, err := primitive.ObjectIDFromHex(category.ID)
		if err != nil {
			return nil, 0, err
		}
		categoryMap[categoryID] = category.Name
	}

	// Relatórios separados por tipo (income/expense)
	reportMap := map[string]map[primitive.ObjectID]*models.Report{
		"income":  {},
		"expense": {},
	}

	// Saldo total
	var totalBalance float64

	for _, transaction := range transactions {
		// Obter o nome da categoria
		categoryName := "Unknown"
		if name, exists := categoryMap[transaction.CategoryID]; exists {
			categoryName = name
		}

		if _, exists := reportMap[transaction.Type][transaction.CategoryID]; !exists {
			reportMap[transaction.Type][transaction.CategoryID] = &models.Report{
				CategoryName: categoryName,
				TotalAmount:  0,
				Transactions: []models.Transaction{},
			}
		}

		report := reportMap[transaction.Type][transaction.CategoryID]
		report.TotalAmount += transaction.Amount
		report.Transactions = append(report.Transactions, transaction)

		if transaction.Type == "income" {
			totalBalance += transaction.Amount
		} else if transaction.Type == "expense" {
			totalBalance -= transaction.Amount
		}
	}

	finalReports := make(map[string][]models.Report)
	for reportType, categories := range reportMap {
		for _, report := range categories {
			finalReports[reportType] = append(finalReports[reportType], *report)
		}
	}

	return finalReports, totalBalance, nil
}
func GeneratePDFReport(reportsByType map[string][]models.Report, totalBalance float64, filePath string, startDate time.Time, endDate time.Time) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetFont("Arial", "", 12)

	// Adicionar título ao relatório com cor
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.SetTextColor(0, 0, 0) // Cor preta para o título
	pdf.Cell(0, 10, "Financial Report")
	pdf.Ln(12)

	// Adicionar intervalo de datas com cor diferente
	pdf.SetFont("Arial", "", 12)
	pdf.SetTextColor(0, 0, 255) // Azul para as datas
	pdf.Cell(0, 10, fmt.Sprintf("Start Date: %s", startDate.Format("2006-01-02")))
	pdf.Ln(6)
	pdf.Cell(0, 10, fmt.Sprintf("End Date: %s", endDate.Format("2006-01-02")))
	pdf.Ln(12)

	// Adicionar seções por tipo (Income e Expense)
	for reportType, reports := range reportsByType {
		// Usar cor para o título da seção
		pdf.SetFont("Arial", "B", 15)
		if reportType == "income" {
			pdf.SetTextColor(0, 128, 0) // Verde para Income
		} else {
			pdf.SetTextColor(255, 0, 0) // Vermelho para Expense
		}
		pdf.Cell(0, 10, strings.Title(reportType))
		pdf.Ln(10)

		pdf.SetFont("Arial", "", 12)
		for _, report := range reports {
			// Exibir a categoria
			pdf.SetTextColor(0, 0, 0) // Preto para a categoria
			pdf.SetFont("Arial", "B", 12)
			pdf.Cell(0, 10, fmt.Sprintf("Category: %s", report.CategoryName))
			pdf.SetFont("Arial", "", 12)
			pdf.Ln(6)

			// Exibir as transações da categoria com cores baseadas no tipo
			for _, transaction := range report.Transactions {
				if transaction.Type == "income" {
					pdf.SetTextColor(0, 128, 0) // Verde para Income
				} else {
					pdf.SetTextColor(255, 0, 0) // Vermelho para Expense
				}
				pdf.Cell(0, 10, fmt.Sprintf("- %s: %.2f", transaction.Description, transaction.Amount))
				pdf.Ln(6)
			}

			// Exibir o total da categoria
			pdf.SetTextColor(0, 0, 0) // Preto
			pdf.Cell(0, 10, fmt.Sprintf("Total for %s: %.2f", report.CategoryName, report.TotalAmount))
			pdf.Ln(10)
		}

		pdf.Ln(10)
	}

	pdf.SetFont("Arial", "B", 14)
	pdf.SetTextColor(255, 165, 0)
	pdf.Cell(0, 10, fmt.Sprintf("Total Balance: %.2f", totalBalance))
	pdf.Ln(10)

	err := pdf.OutputFileAndClose(filePath)
	if err != nil {
		return fmt.Errorf("could not save PDF: %w", err)
	}

	return nil
}

func formatFloat(value float64) string {
	return fmt.Sprintf("%.2f", value)
}

func GenerateAndExportPDFReport(startDate, endDate time.Time, filePath string) error {
	// Obter os relatórios categorizados por tipo e o saldo total
	reportsByType, totalBalance, err := GenerateReportByType(startDate, endDate)
	if err != nil {
		return err
	}

	// Gerar o PDF com os relatórios categorizados por tipo
	return GeneratePDFReport(reportsByType, totalBalance, filePath, startDate, endDate)
}
