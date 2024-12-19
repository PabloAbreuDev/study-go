package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID          string             `json:"id" bson:"_id,omitempty"`
	Type        string             `json:"type"`
	Amount      float64            `json:"amount"`
	Description string             `json:"description"`
	CategoryID  primitive.ObjectID `json:"category_id" bson:"category_id"`
	Date        time.Time          `json:"date" bson:"date"`
}

type Date time.Time

func (d *Date) UnmarshalJSON(b []byte) error {
	str := string(b)
	str = str[1 : len(str)-1]
	parsedTime, err := time.Parse("2006-01-02", str)
	if err != nil {
		return err
	}
	*d = Date(parsedTime)
	return nil
}
