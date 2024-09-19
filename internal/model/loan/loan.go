package loanModel

import "time"

type Loan struct {
	ID           string    `json:"id" db:"id"`
	CustomerID   string    `json:"customer_id" db:"customer_id"`
	Name         string    `json:"name" db:"name"`
	Period       int       `json:"period" db:"period"`
	Amount       float64   `json:"amount" db:"amount"`
	InterestRate float64   `json:"interest_rate" db:"interest_rate"`
	TotalAmount  float64   `json:"total_amount" db:"total_amount"`
	Status       string    `json:"status" db:"status"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}
