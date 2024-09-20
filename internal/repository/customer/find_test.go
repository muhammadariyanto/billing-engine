package customerRepository

import (
	"context"
	"github.com/google/uuid"
	customerModel "github.com/muhammadariyanto/billing-engine/internal/model/customer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindByID(t *testing.T) {
	validCustomerID := uuid.NewString()

	testCases := []struct {
		name    string
		wantErr bool
		repo    *customerRepository
	}{
		{
			name:    "ERROR: Customer data is not found",
			wantErr: true,
			repo:    &customerRepository{},
		},
		{
			name:    "SUCCESS: Fetch data",
			wantErr: false,
			repo: &customerRepository{
				data: map[string]*customerModel.Customer{
					validCustomerID: {ID: validCustomerID, Name: "John Doe"},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			customer, err := tc.repo.FindByID(context.Background(), validCustomerID)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, customer)
			}
		})
	}

}
