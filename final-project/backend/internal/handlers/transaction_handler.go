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

	result := repositories.CreateTransaction(transaction)

	c.JSON(http.StatusOK, gin.H{"result": result})
}
