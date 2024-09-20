package customerService_test

import (
	"context"
	"errors"
	"github.com/google/uuid"
	billingModel "github.com/muhammadariyanto/billing-engine/internal/model/billing"
	loanModel "github.com/muhammadariyanto/billing-engine/internal/model/loan"
	. "github.com/muhammadariyanto/billing-engine/internal/service/customer"
	repositoryMock "github.com/muhammadariyanto/billing-engine/mocks/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestIsDelinquent(t *testing.T) {
	customerRepo := repositoryMock.NewICustomerRepository(t)
	loanRepo := repositoryMock.NewILoanRepository(t)
	billingRepo := repositoryMock.NewIBillingRepository(t)
	svc := New(customerRepo, loanRepo, billingRepo)

	testCases := []struct {
		name                 string
		wantErr              bool
		setupMock            func()
		expectedIsDelinquent bool
	}{
		{
			name:    "ERROR: FetchUncompletedByCustomerID error",
			wantErr: true,
			setupMock: func() {
				loanRepo.On("FetchUncompletedByCustomerID",
					mock.Anything,
					mock.AnythingOfType("string"),
				).Once().Return(nil, errors.New("some error"))
			},
		},
		{
			name:    "SUCCESS: Empty loan",
			wantErr: false,
			setupMock: func() {
				loanRepo.On("FetchUncompletedByCustomerID",
					mock.Anything,
					mock.AnythingOfType("string"),
				).Once().Return([]*loanModel.Loan{}, nil)
			},
		},
		{
			name:    "SUCCESS: With error fetching unpaid billing",
			wantErr: false,
			setupMock: func() {
				loanRepo.On("FetchUncompletedByCustomerID",
					mock.Anything,
					mock.AnythingOfType("string"),
				).Return([]*loanModel.Loan{
					{
						ID: uuid.NewString(),
					},
				}, nil)

				billingRepo.On("FetchUnpaidByLoanID",
					mock.Anything,
					mock.AnythingOfType("string"),
				).Once().Return(nil, errors.New("error when fetch data"))
			},
		},
		{
			name:    "SUCCESS: Customer is delinquent",
			wantErr: false,
			setupMock: func() {
				billingRepo.On("FetchUnpaidByLoanID",
					mock.Anything,
					mock.AnythingOfType("string"),
				).Once().Return([]*billingModel.Billing{
					{
						DueDate: time.Now().Add(-4 * 7 * 24 * time.Hour),
					},
					{
						DueDate: time.Now().Add(-3 * 7 * 24 * time.Hour),
					},
				}, nil)
			},
			expectedIsDelinquent: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock()

			isDelinquent, err := svc.IsDelinquent(context.Background(), uuid.NewString())
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedIsDelinquent, isDelinquent)
			}

			customerRepo.AssertExpectations(t)
			billingRepo.AssertExpectations(t)
			loanRepo.AssertExpectations(t)
		})
	}
}
