package loanRepository

import (
	"context"
	"github.com/google/uuid"
	loanModel "github.com/muhammadariyanto/billing-engine/internal/model/loan"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert(t *testing.T) {
	loan := &loanModel.Loan{
		ID:   uuid.NewString(),
		Name: "Loan 1",
	}

	testCases := []struct {
		name    string
		wantErr bool
		repo    *loanRepository
		loan    *loanModel.Loan
	}{
		{
			name:    "ERROR: loan ID already existed",
			wantErr: true,
			repo: &loanRepository{
				data: map[string]*loanModel.Loan{
					loan.ID: loan,
				},
			},
			loan: loan,
		},
		{
			name:    "SUCCESS: Insert data",
			wantErr: false,
			repo: &loanRepository{
				data: make(map[string]*loanModel.Loan),
			},
			loan: loan,
		},
		{
			name:    "SUCCESS: Insert data on empty ID",
			wantErr: false,
			repo: &loanRepository{
				data: make(map[string]*loanModel.Loan),
			},
			loan: &loanModel.Loan{
				ID:   "",
				Name: "Loan 2",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.repo.Insert(context.Background(), tc.loan)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}
