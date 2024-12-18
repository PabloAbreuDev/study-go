package handlers

import (
	"example/backend/internal/models"
	"example/backend/internal/repositories"
	"example/backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTransactions(c *gin.Context) {

	transactions := services.FetchTransactions()
	c.JSON(http.StatusOK, gin.H{"data": transactions})
}

func PostTransaction(c *gin.Context) {
	var transaction models.Transaction

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	result := services.PostTransactions(transaction)

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func GetTransactionById(c *gin.Context) {
	transactionId := c.Param("id")

	transaction, err := services.GetTransactionByID(transactionId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": transaction})
}

func UpdateTransaction(c *gin.Context) {
	id := c.Param("id") // Get the transaction ID from the URL parameter

	var updatedData models.Transaction
	// Bind the request body to the updatedData struct
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	// Call the repository function to update the transaction
	err := repositories.UpdateTransaction(id, updatedData)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction updated successfully"})
}

func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")

	err := services.DeleteTransaction(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete transaction: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}
