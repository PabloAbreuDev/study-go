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

func GetAllBudgets() []models.Budget {
	collection := db.GetCollection("finances-app", "budgets")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer cursor.Close(ctx)

	var budgets []models.Budget
	for cursor.Next(ctx) {
		var budget models.Budget
		if err = cursor.Decode(&budget); err != nil {
			log.Fatal(err)
			return nil
		}
		budgets = append(budgets, budget)
	}

	return budgets
}

func CreateBudget(budget models.Budget) interface{} {
	collection := db.GetCollection("finances-app", "budgets")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, budget)
	if err != nil {
		log.Fatal(err)
	}

	return result.InsertedID
}

func GetBudgetByID(id string) (*models.Budget, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	collection := db.GetCollection("finances-app", "budgets")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var budget models.Budget
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&budget)
	if err != nil {
		return nil, err
	}

	return &budget, nil
}

func UpdateBudget(id string, updatedData models.Budget) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := db.GetCollection("finances-app", "budgets")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"name":       updatedData.Name,
			"amount":     updatedData.Amount,
			"start_date": updatedData.StartDate,
			"end_date":   updatedData.EndDate,
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	return err
}
