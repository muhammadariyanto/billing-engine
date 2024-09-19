package customerService

import (
	"context"
	"fmt"
	"time"
)

const MaxAllowedDueBilling = 2

func (s *customerService) IsDelinquent(ctx context.Context, customerID string) (bool, error) {
	// Fetch uncompleted loan
	loans, err := s.loanRepository.FetchUncompletedByCustomerID(ctx, customerID)
	if err != nil {
		return false, err
	}

	if len(loans) == 0 {
		return false, nil
	}

	countDueBilling := 0
	for _, loan := range loans {
		// fetch unpaid billing
		billings, err := s.billingRepository.FetchUnpaidByLoanID(ctx, loan.ID)
		if err != nil {
			continue
		}

		for _, billing := range billings {
			fmt.Println(billing.DueDate, time.Now())

			if billing.DueDate.Before(time.Now()) {
				countDueBilling += 1
			}

			if countDueBilling > MaxAllowedDueBilling {
				return true, nil
			}
		}
	}

	return false, nil
}
