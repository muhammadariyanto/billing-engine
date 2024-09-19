package loanRepository

import (
	"context"
	"errors"
	"github.com/muhammadariyanto/billing-engine/constant"
	loanModel "github.com/muhammadariyanto/billing-engine/internal/model/loan"
)

func (r *loanRepository) FetchUncompletedByCustomerID(ctx context.Context, customerID string) ([]*loanModel.Loan, error) {
	resp := make([]*loanModel.Loan, 0)
	for _, loan := range r.data {
		if loan.CustomerID == customerID && loan.Status != constant.LoanStatusCompleted {
			resp = append(resp, loan)
		}
	}

	if len(resp) == 0 {
		return nil, errors.New("data is not found")
	}

	return resp, nil
}
