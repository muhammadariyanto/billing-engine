package billingModel

import "time"

type Billing struct {
	ID          string     `json:"id" db:"id"`
	LoanID      string     `json:"loan_id" db:"loan_id"`
	Sequence    int        `json:"sequence" db:"sequence"`
	Date        time.Time  `json:"date" db:"date"`
	DueDate     time.Time  `json:"due_date" db:"due_date"`
	PaymentDate *time.Time `json:"payment_date" db:"payment_date"`
	Amount      float64    `json:"amount" db:"amount"`
	Interest    float64    `json:"interest" db:"interest"`
	TotalAmount float64    `json:"total_amount" db:"total_amount"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
}
