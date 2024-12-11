package handlers

import (
	"example/backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTransactions(c *gin.Context) {

	transactions := services.FetchTransactions()
	c.JSON(http.StatusOK, gin.H{"data": transactions})
}
