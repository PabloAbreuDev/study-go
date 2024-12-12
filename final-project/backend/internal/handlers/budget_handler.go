package handlers

import (
	"example/backend/internal/models"
	"example/backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBudgets(c *gin.Context) {
	budgets := services.FetchBudgets()
	c.JSON(http.StatusOK, gin.H{"data": budgets})
}

func PostBudget(c *gin.Context) {
	var budget models.Budget

	if err := c.ShouldBindJSON(&budget); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	result := services.PostBudget(budget)
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func GetBudgetByID(c *gin.Context) {
	id := c.Param("id")
	budget, err := services.GetBudgetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Budget not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": budget})
}

func UpdateBudget(c *gin.Context) {
	id := c.Param("id")
	var updatedBudget models.Budget

	if err := c.ShouldBindJSON(&updatedBudget); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	err := services.UpdateBudget(id, updatedBudget)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update budget"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Budget updated successfully"})
}
