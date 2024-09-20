package billingModel

import "time"

type MakePaymentRequest struct {
	LoanID        string    `json:"loan_id" validate:"required"`
	PaymentAmount float64   `json:"payment_amount" validate:"min=0"`
	PaymentDate   time.Time `json:"payment_date" validate:"required"`
}
