package repositories

import (
	"context"
	"example/backend/db"
	"example/backend/internal/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func GetTransactionByID(id string) (*models.Transaction, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	collection := db.GetCollection("finances-app", "transactions")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var transaction models.Transaction
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&transaction)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func UpdateTransaction(id string, updatedData models.Transaction) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := db.GetCollection("finances-app", "transactions")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"type":        updatedData.Type,
			"amount":      updatedData.Amount,
			"description": updatedData.Description,
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	return err
}

func DeleteTransaction(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := db.GetCollection("finances-app", "transactions")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}
