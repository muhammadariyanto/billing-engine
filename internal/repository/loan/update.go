package loanRepository

import (
	"context"
	"errors"
	loanModel "github.com/muhammadariyanto/billing-engine/internal/model/loan"
)

func (r *loanRepository) Update(ctx context.Context, loan *loanModel.Loan) error {
	if _, exists := r.data[loan.ID]; !exists {
		return errors.New("loan is not exists")
	}

	r.data[loan.ID] = loan
	return nil
}
