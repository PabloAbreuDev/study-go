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

func GetAllCategories() []models.Category {
	collection := db.GetCollection("finances-app", "categories")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer cursor.Close(ctx)

	var categories []models.Category
	for cursor.Next(ctx) {
		var category models.Category
		if err = cursor.Decode(&category); err != nil {
			log.Fatal(err)
			return nil
		}
		categories = append(categories, category)
	}

	return categories
}

func CreateCategory(category models.Category) interface{} {
	collection := db.GetCollection("finances-app", "categories")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, category)
	if err != nil {
		log.Fatal(err)
	}

	return result.InsertedID
}

func GetCategoryByID(id string) (*models.Category, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	collection := db.GetCollection("finances-app", "categories")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var category models.Category
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&category)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func UpdateCategory(id string, updatedData models.Category) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := db.GetCollection("finances-app", "categories")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"name":        updatedData.Name,
			"description": updatedData.Description,
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	return err
}
