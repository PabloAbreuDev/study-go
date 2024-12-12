package models

type Transaction struct {
	ID          string  `json:"id" bson:"_id,omitempty"`
	Type        string  `json:"type"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}
