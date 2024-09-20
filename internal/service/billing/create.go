package billingService

import (
	"context"
	billingModel "github.com/muhammadariyanto/billing-engine/internal/model/billing"
	"sync"
	"time"
)

func (s *billingService) CreateSchedule(ctx context.Context, loanID string, startDate time.Time) error {
	// Find loan first
	loan, err := s.loanRepository.FindByID(ctx, loanID)
	if err != nil {
		return err
	}

	// Calculate amount and interest value per period
	amount := loan.Amount / float64(loan.Period)
	interest := amount * loan.InterestRate

	// Use goroutine instead
	wg := sync.WaitGroup{}
	wg.Add(loan.Period)

	// Generate billing schedule by loan period (assume period in weekly)
	for i := 0; i < loan.Period; i++ {
		date := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, startDate.Location())

		// We assume only apply weekly period
		weeklyDuration := 7 * 24 * time.Hour
		dueDate := date.Add(weeklyDuration - time.Millisecond)

		billing := &billingModel.Billing{
			LoanID:      loanID,
			Sequence:    i + 1,
			Date:        date,
			DueDate:     dueDate,
			Amount:      amount,
			Interest:    interest,
			TotalAmount: amount + interest,
		}

		// Insert billing schedule, skip on error, it should be use db transaction in real use.
		go func() {
			defer wg.Done()

			_ = s.billingRepository.Insert(ctx, billing)
		}()

		// Set startDate with a week forward
		startDate = startDate.Add(weeklyDuration)
	}

	wg.Wait()

	return nil
}
