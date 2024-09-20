package billingRepository

import (
	"context"
	"github.com/google/uuid"
	billingModel "github.com/muhammadariyanto/billing-engine/internal/model/billing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdate(t *testing.T) {
	billing := &billingModel.Billing{
		ID: uuid.NewString(),
	}

	testCases := []struct {
		name    string
		wantErr bool
		repo    *billingRepository
	}{
		{
			name:    "ERROR: Billing ID already existed",
			wantErr: true,
			repo: &billingRepository{
				data: make(map[string]*billingModel.Billing),
			},
		},
		{
			name:    "SUCCESS: Update data",
			wantErr: false,
			repo: &billingRepository{
				data: map[string]*billingModel.Billing{
					billing.ID: billing,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.repo.Update(context.Background(), billing)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}
