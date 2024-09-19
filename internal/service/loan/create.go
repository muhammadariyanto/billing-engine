package loanService

import (
	"context"
	loanModel "github.com/muhammadariyanto/billing-engine/internal/model/loan"
)

func (s *loanService) CreateLoan(ctx context.Context, loan *loanModel.Loan) error {
	// Customer should be existed
	if _, err := s.customerRepository.FindByID(ctx, loan.CustomerID); err != nil {
		return err
	}

	return s.loanRepository.Insert(ctx, loan)
}
