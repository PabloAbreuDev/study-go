package routes

import (
	"example/backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/transactions", handlers.GetTransactions)
	r.POST("/transactions", handlers.PostTransaction)

	return r
}
