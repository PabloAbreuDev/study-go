package models

import "time"

type Transaction struct {
	ID          string
	Type        string
	Amount      float64
	Description string
	Category    string
	Date        time.Time
}
