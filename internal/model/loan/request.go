package loanModel

import "time"

type ApplyLoanRequest struct {
	CustomerID   string    `json:"customer_id" validate:"required"`
	Name         string    `json:"name" validate:"required"`
	Period       int       `json:"period" validate:"min=1"`
	Amount       float64   `json:"amount" validate:"min=0"`
	InterestRate float64   `json:"interest_rate" validate:"min=0.0,max=1.0"`
	StartDate    time.Time `json:"start_date" validate:"required"`
}
