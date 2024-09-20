package billingService

import (
	"context"
	billingModel "github.com/muhammadariyanto/billing-engine/internal/model/billing"
)

func (s *billingService) FetchAllByLoanID(ctx context.Context, loanID string) ([]*billingModel.Billing, error) {
	return s.billingRepository.FetchAllByLoanID(ctx, loanID)
}
