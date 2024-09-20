package customerService_test

import (
	"context"
	customerModel "github.com/muhammadariyanto/billing-engine/internal/model/customer"
	. "github.com/muhammadariyanto/billing-engine/internal/service/customer"
	repositoryMock "github.com/muhammadariyanto/billing-engine/mocks/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestRegister(t *testing.T) {
	customerRepo := repositoryMock.NewICustomerRepository(t)
	svc := New(customerRepo, nil, nil)

	testCases := []struct {
		name      string
		wantErr   bool
		setupMock func()
	}{
		{
			name:    "SUCCESS",
			wantErr: false,
			setupMock: func() {
				customerRepo.On("Insert",
					mock.Anything,
					mock.AnythingOfType("*customerModel.Customer"),
				).Once().Return(nil)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock()

			err := svc.Register(context.Background(), &customerModel.Customer{})
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			customerRepo.AssertExpectations(t)
		})
	}
}
