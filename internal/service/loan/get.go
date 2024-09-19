package loanService

import (
	"context"
	"github.com/muhammadariyanto/billing-engine/constant"
)

func (s *loanService) GetOutstanding(ctx context.Context, loanID string) (float64, error) {
	// Find loan first
	loan, err := s.loanRepository.FindByID(ctx, loanID)
	if err != nil {
		return 0, err
	}

	// Return 0 when loan status completed
	if loan.Status == constant.LoanStatusCompleted {
		return 0, nil
	}

	return s.billingRepository.SumUnpaidByLoanID(ctx, loanID), nil
}
