package loanRepository

import (
	"context"
	"github.com/google/uuid"
	loanModel "github.com/muhammadariyanto/billing-engine/internal/model/loan"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdate(t *testing.T) {
	loan := &loanModel.Loan{
		ID: uuid.NewString(),
	}

	testCases := []struct {
		name    string
		wantErr bool
		repo    *loanRepository
	}{
		{
			name:    "ERROR: loan ID already existed",
			wantErr: true,
			repo: &loanRepository{
				data: make(map[string]*loanModel.Loan),
			},
		},
		{
			name:    "SUCCESS: Update data",
			wantErr: false,
			repo: &loanRepository{
				data: map[string]*loanModel.Loan{
					loan.ID: loan,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.repo.Update(context.Background(), loan)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}
