package customerRepository

import (
	"context"
	"github.com/google/uuid"
	customerModel "github.com/muhammadariyanto/billing-engine/internal/model/customer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert(t *testing.T) {
	customer := &customerModel.Customer{
		ID:   uuid.NewString(),
		Name: "John",
	}

	testCases := []struct {
		name     string
		wantErr  bool
		repo     *customerRepository
		customer *customerModel.Customer
	}{
		{
			name:    "ERROR: Customer ID already existed",
			wantErr: true,
			repo: &customerRepository{
				data: map[string]*customerModel.Customer{
					customer.ID: customer,
				},
			},
			customer: customer,
		},
		{
			name:    "SUCCESS: Insert data",
			wantErr: false,
			repo: &customerRepository{
				data: make(map[string]*customerModel.Customer),
			},
			customer: customer,
		},
		{
			name:    "SUCCESS: Insert data on empty ID",
			wantErr: false,
			repo: &customerRepository{
				data: make(map[string]*customerModel.Customer),
			},
			customer: &customerModel.Customer{
				ID:   "",
				Name: "John",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.repo.Insert(context.Background(), tc.customer)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}
