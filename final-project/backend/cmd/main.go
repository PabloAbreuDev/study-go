package main

import (
	"example/backend/internal/routes"
	"log"
)

// type Transaction struct {
// 	ID          string
// 	Type        string
// 	Amount      float64
// 	Description string
// 	Category    string
// 	Date        time.Time
// }

// var transactions = []Transaction{
// 	{
// 		ID:          "12345",
// 		Type:        "despesa",
// 		Amount:      150.75,
// 		Description: "Compra de supermercado",
// 		Category:    "Alimentação",
// 		Date:        time.Now(),
// 	},
// 	{
// 		ID:          "12346",
// 		Type:        "despesa2",
// 		Amount:      10.75,
// 		Description: "Compra de supermercado",
// 		Category:    "Alimentação",
// 		Date:        time.Now(),
// 	},
// }

// func getTransactions(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, transactions)
// }

// func postTransactions(c *gin.Context) {
// 	var newTransaction Transaction

// 	if err := c.BindJSON(&newTransaction); err != nil {
// 		return
// 	}
// 	transactions = append(transactions, newTransaction)
// 	c.IndentedJSON(http.StatusCreated, newTransaction)
// }

// func getTransactionById(c *gin.Context) {
// 	id := c.Param("id")

// 	for _, a := range transactions {
// 		if a.ID == id {
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "transaction not found"})
// }

func main() {

	// router := gin.Default()
	// router.GET("/transactions", getTransactions)
	// router.GET("/transactions/:id", getTransactionById)
	// router.POST("/transactions", postTransactions)
	// router.Run("localhost:8080")

	r := routes.SetupRoutes()
	log.Println("Server running on http://localhost:8080")
	log.Fatal(r.Run(":8080"))

}
