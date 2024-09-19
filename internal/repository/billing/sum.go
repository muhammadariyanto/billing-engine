package billingRepository

import (
	"context"
)

func (r *billingRepository) SumUnpaidByLoanID(ctx context.Context, loanID string) float64 {
	sumUnpaid := 0.00

	for _, billing := range r.data {
		if billing.LoanID == loanID && billing.PaymentDate == nil {
			sumUnpaid += billing.TotalAmount
		}
	}

	return sumUnpaid
}
