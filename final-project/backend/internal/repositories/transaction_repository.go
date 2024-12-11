package repositories

import (
	"context"
	"example/backend/db"
	"example/backend/internal/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllTransactions() []models.Transaction {
	collection := db.GetCollection("finances-app", "transactions")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer cursor.Close(ctx)
	var transactions []models.Transaction

	for cursor.Next(ctx) {
		var transaction models.Transaction
		if err = cursor.Decode(&transaction); err != nil {
			log.Fatal(err)
			return nil
		}
		transactions = append(transactions, transaction)
	}

	return transactions
}

func CreateTransaction(transaction models.Transaction) interface{} {
	collection := db.GetCollection("finances-app", "transactions")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, transaction)
	if err != nil {
		log.Fatal(err)
	}

	return result.InsertedID
}
