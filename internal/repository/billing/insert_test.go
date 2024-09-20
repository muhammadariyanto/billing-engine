package billingRepository

import (
	"context"
	"github.com/google/uuid"
	billingModel "github.com/muhammadariyanto/billing-engine/internal/model/billing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert(t *testing.T) {
	billing := &billingModel.Billing{
		ID:     uuid.NewString(),
		LoanID: uuid.NewString(),
	}

	testCases := []struct {
		name    string
		wantErr bool
		repo    *billingRepository
		billing *billingModel.Billing
	}{
		{
			name:    "ERROR: Billing ID already existed",
			wantErr: true,
			repo: &billingRepository{
				data: map[string]*billingModel.Billing{
					billing.ID: billing,
				},
			},
			billing: billing,
		},
		{
			name:    "SUCCESS: Insert data",
			wantErr: false,
			repo: &billingRepository{
				data: make(map[string]*billingModel.Billing),
			},
			billing: billing,
		},
		{
			name:    "SUCCESS: Insert data on empty ID",
			wantErr: false,
			repo: &billingRepository{
				data: make(map[string]*billingModel.Billing),
			},
			billing: &billingModel.Billing{
				ID:     "",
				LoanID: uuid.NewString(),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.repo.Insert(context.Background(), tc.billing)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}
