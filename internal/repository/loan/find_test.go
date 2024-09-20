package loanRepository

import (
	"context"
	"github.com/google/uuid"
	loanModel "github.com/muhammadariyanto/billing-engine/internal/model/loan"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindByID(t *testing.T) {
	validLoanID := uuid.NewString()

	testCases := []struct {
		name    string
		wantErr bool
		repo    *loanRepository
	}{
		{
			name:    "ERROR: loan data is not found",
			wantErr: true,
			repo:    &loanRepository{},
		},
		{
			name:    "SUCCESS: Fetch data",
			wantErr: false,
			repo: &loanRepository{
				data: map[string]*loanModel.Loan{
					validLoanID: {ID: validLoanID, Name: "Loan 1"},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			loan, err := tc.repo.FindByID(context.Background(), validLoanID)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, loan)
			}
		})
	}

}
