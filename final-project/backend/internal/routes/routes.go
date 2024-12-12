package routes

import (
	"example/backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Transaction routes
	r.GET("/transactions", handlers.GetTransactions)
	r.GET("/transactions/:id", handlers.GetTransactionById)
	r.POST("/transactions", handlers.PostTransaction)
	r.PUT("/transactions/:id", handlers.UpdateTransaction)

	// Category routes
	r.GET("/categories", handlers.GetCategories)
	r.POST("/categories", handlers.PostCategory)
	r.GET("/categories/:id", handlers.GetCategoryByID)
	r.PUT("/categories/:id", handlers.UpdateCategory)

	// Budget routes
	r.GET("/budgets", handlers.GetBudgets)
	r.POST("/budgets", handlers.PostBudget)
	r.GET("/budgets/:id", handlers.GetBudgetByID)
	r.PUT("/budgets/:id", handlers.UpdateBudget)

	return r
}
