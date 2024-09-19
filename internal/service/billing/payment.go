package billingService

import (
	"context"
	"fmt"
	"github.com/muhammadariyanto/billing-engine/constant"
	billingModel "github.com/muhammadariyanto/billing-engine/internal/model/billing"
	"time"
)

func (s *billingService) MakePayment(ctx context.Context, loanID string, paymentAmount float64, paymentDate time.Time) error {
	// Find loan first
	loan, err := s.loanRepository.FindByID(ctx, loanID)
	if err != nil {
		return err
	}

	// Return if loan is not IN_PROGRESS status
	if loan.Status != constant.LoanStatusInProgress {
		return fmt.Errorf("cannot proceed: loan with ID %s has already been completed or is no longer in progress", loan.ID)
	}

	// Fetch all schedule (sort by sequence)
	billings, err := s.billingRepository.FetchAllByLoanID(ctx, loanID)
	if err != nil {
		return err
	}

	// Select billing to pay
	billingToPay := &billingModel.Billing{}
	for _, billing := range billings {
		if billing.PaymentDate == nil {
			billingToPay = billing
			break
		}
	}

	// Return if paymentAmount is not match
	if billingToPay.TotalAmount != paymentAmount {
		return fmt.Errorf("payment amount mismatch: expected %.2f but received %.2f for billing ID %s", billingToPay.TotalAmount, paymentAmount, billingToPay.ID)
	}

	// Set payment date
	billingToPay.PaymentDate = &paymentDate

	// Update billing
	return s.billingRepository.Update(ctx, billingToPay)
}
