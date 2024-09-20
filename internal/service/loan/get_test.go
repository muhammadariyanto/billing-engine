package loanService_test

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/muhammadariyanto/billing-engine/constant"
	loanModel "github.com/muhammadariyanto/billing-engine/internal/model/loan"
	. "github.com/muhammadariyanto/billing-engine/internal/service/loan"
	repositoryMock "github.com/muhammadariyanto/billing-engine/mocks/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestGetOutstanding(t *testing.T) {
	billingRepo := repositoryMock.NewIBillingRepository(t)
	loanRepo := repositoryMock.NewILoanRepository(t)
	svc := New(loanRepo, nil, billingRepo)

	testCases := []struct {
		name                string
		wantErr             bool
		setupMock           func()
		expectedOutstanding float64
	}{
		{
			name:    "ERROR: Find loan error",
			wantErr: true,
			setupMock: func() {
				loanRepo.On("FindByID",
					mock.Anything,
					mock.AnythingOfType("string"),
				).Once().Return(nil, errors.New("some error"))
			},
		},
		{
			name:    "SUCCESS: Loan status is completed",
			wantErr: false,
			setupMock: func() {
				loanRepo.On("FindByID",
					mock.Anything,
					mock.AnythingOfType("string"),
				).Once().Return(&loanModel.Loan{
					Status: constant.LoanStatusCompleted,
				}, nil)
			},
			expectedOutstanding: 0,
		},
		{
			name:    "SUCCESS: Loan status is completed",
			wantErr: false,
			setupMock: func() {
				loanRepo.On("FindByID",
					mock.Anything,
					mock.AnythingOfType("string"),
				).Once().Return(&loanModel.Loan{
					Status: constant.LoanStatusInProgress,
				}, nil)

				billingRepo.On("SumUnpaidByLoanID",
					mock.Anything,
					mock.AnythingOfType("string"),
				).Once().Return(2_500_000.00)
			},
			expectedOutstanding: 2_500_000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock()

			outstanding, err := svc.GetOutstanding(context.Background(), uuid.NewString())
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOutstanding, outstanding)
			}

			billingRepo.AssertExpectations(t)
			loanRepo.AssertExpectations(t)
		})
	}
}
