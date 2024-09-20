package billingRepository

import (
	"context"
	"errors"
	billingModel "github.com/muhammadariyanto/billing-engine/internal/model/billing"
	"sort"
)

func (r *billingRepository) FetchAllByLoanID(ctx context.Context, loanID string) ([]*billingModel.Billing, error) {
	resp := make([]*billingModel.Billing, 0)
	for _, billing := range r.data {
		if billing.LoanID == loanID {
			resp = append(resp, billing)
		}
	}

	if len(resp) == 0 {
		return nil, errors.New("billing data is not found")
	}

	// Order by sequence
	sort.Slice(resp, func(i, j int) bool {
		return resp[i].Sequence < resp[j].Sequence
	})

	return resp, nil
}

func (r *billingRepository) FetchUnpaidByLoanID(ctx context.Context, loanID string) ([]*billingModel.Billing, error) {
	resp := make([]*billingModel.Billing, 0)
	for _, billing := range r.data {
		if billing.LoanID == loanID && billing.PaymentDate == nil {
			resp = append(resp, billing)
		}
	}

	if len(resp) == 0 {
		return nil, errors.New("loan data is not found")
	}

	// Order by sequence
	sort.Slice(resp, func(i, j int) bool {
		return resp[i].Sequence < resp[j].Sequence
	})

	return resp, nil
}
