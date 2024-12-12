package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Budget struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name"`
	Amount    float64            `json:"amount"`
	StartDate string             `json:"start_date"`
	EndDate   string             `json:"end_date"`
}
