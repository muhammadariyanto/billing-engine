package loanRepository

import (
	"context"
	"github.com/google/uuid"
	"github.com/muhammadariyanto/billing-engine/constant"
	loanModel "github.com/muhammadariyanto/billing-engine/internal/model/loan"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFetchUncompletedByCustomerID(t *testing.T) {
	customerID := uuid.NewString()

	testCases := []struct {
		name    string
		wantErr bool
		repo    *loanRepository
	}{
		{
			name:    "ERROR: Loans data is not found",
			wantErr: true,
			repo:    &loanRepository{},
		},
		{
			name:    "SUCCESS: Fetch data",
			wantErr: false,
			repo: &loanRepository{
				data: map[string]*loanModel.Loan{
					"1": {CustomerID: customerID, Status: constant.LoanStatusInProgress},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			loans, err := tc.repo.FetchUncompletedByCustomerID(context.Background(), customerID)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, loans)
			}
		})
	}
}
