package billingRepository

import (
	"context"
	"github.com/google/uuid"
	billingModel "github.com/muhammadariyanto/billing-engine/internal/model/billing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFetchAllByLoanID(t *testing.T) {
	validLoanID := uuid.NewString()

	testCases := []struct {
		name    string
		wantErr bool
		repo    *billingRepository
	}{
		{
			name:    "ERROR: Billings data is not found",
			wantErr: true,
			repo:    &billingRepository{},
		},
		{
			name:    "SUCCESS: Fetch data",
			wantErr: false,
			repo: &billingRepository{
				data: map[string]*billingModel.Billing{
					"1": {LoanID: validLoanID, Sequence: 1},
					"2": {LoanID: validLoanID, Sequence: 2},
					"3": {LoanID: validLoanID, Sequence: 3},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			billings, err := tc.repo.FetchAllByLoanID(context.Background(), validLoanID)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, billings)
			}
		})
	}

}

func TestFetchUnpaidByLoanID(t *testing.T) {
	validLoanID := uuid.NewString()

	testCases := []struct {
		name    string
		wantErr bool
		repo    *billingRepository
	}{
		{
			name:    "ERROR: Billings data is not found",
			wantErr: true,
			repo:    &billingRepository{},
		},
		{
			name:    "SUCCESS: Fetch data",
			wantErr: false,
			repo: &billingRepository{
				data: map[string]*billingModel.Billing{
					"1": {LoanID: validLoanID, Sequence: 1},
					"2": {LoanID: validLoanID, Sequence: 2},
					"3": {LoanID: validLoanID, Sequence: 3},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			billings, err := tc.repo.FetchUnpaidByLoanID(context.Background(), validLoanID)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, billings)
			}
		})
	}

}
