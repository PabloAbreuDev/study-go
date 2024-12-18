package models

type Report struct {
	CategoryName string        `json:"category_name"`
	TotalAmount  float64       `json:"total_amount"`
	Transactions []Transaction `json:"transactions"`
}
