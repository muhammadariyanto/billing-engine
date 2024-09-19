package loanRepository

import (
	"context"
	"errors"
	loanModel "github.com/muhammadariyanto/billing-engine/internal/model/loan"
)

func (r *loanRepository) FindByID(ctx context.Context, loanID string) (*loanModel.Loan, error) {
	loan, ok := r.data[loanID]
	if !ok {
		return nil, errors.New("loan not found")
	}
	return loan, nil
}
