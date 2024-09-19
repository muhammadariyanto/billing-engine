package billingService

import (
	"context"
	"fmt"
	"github.com/muhammadariyanto/billing-engine/constant"
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
	billings, err := s.billingRepository.FetchUnpaidByLoanID(ctx, loanID)
	if err != nil {
		return err
	}

	// Select first billing to pay
	billingToPay := billings[0]

	// Return if paymentAmount is not match
	if billingToPay.TotalAmount != paymentAmount {
		return fmt.Errorf("payment amount mismatch: expected %.2f but received %.2f for billing ID %s", billingToPay.TotalAmount, paymentAmount, billingToPay.ID)
	}

	// Set payment date
	billingToPay.PaymentDate = &paymentDate

	// Update billing
	if err := s.billingRepository.Update(ctx, billingToPay); err != nil {
		return err
	}

	// If last payment then update loan status
	if billingToPay.Sequence == loan.Period {
		loan.Status = constant.LoanStatusCompleted
		if err := s.loanRepository.Update(ctx, loan); err != nil {
			return err
		}
	}

	return nil
}
