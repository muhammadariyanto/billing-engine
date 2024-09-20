package billingRepository

import (
	"context"
	"github.com/google/uuid"
	billingModel "github.com/muhammadariyanto/billing-engine/internal/model/billing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumUnpaidByLoanID(t *testing.T) {
	validLoanID := uuid.NewString()

	testCases := []struct {
		name        string
		wantErr     bool
		repo        *billingRepository
		expectedSum float64
	}{
		{
			name:    "SUCCESS: Sum unpaid billing by loan",
			wantErr: false,
			repo: &billingRepository{
				data: map[string]*billingModel.Billing{
					"1": {LoanID: validLoanID, Sequence: 1, TotalAmount: 500_000},
					"2": {LoanID: validLoanID, Sequence: 2, TotalAmount: 500_000},
					"3": {LoanID: validLoanID, Sequence: 3, TotalAmount: 500_000},
				},
			},
			expectedSum: 1_500_000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sum := tc.repo.SumUnpaidByLoanID(context.Background(), validLoanID)
			assert.Equal(t, tc.expectedSum, sum)
		})
	}

}
