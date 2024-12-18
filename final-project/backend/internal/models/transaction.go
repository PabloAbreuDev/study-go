package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaction struct {
	ID          string             `json:"id" bson:"_id,omitempty"`
	Type        string             `json:"type"`
	Amount      float64            `json:"amount"`
	Description string             `json:"description"`
	CategoryID  primitive.ObjectID `json:"category_id" bson:"category_id"`
}
